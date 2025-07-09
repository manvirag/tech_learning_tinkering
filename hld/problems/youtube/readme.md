## Design youtube  ( this is the source of truth )

### Functional Requirements
1. Ability to upload video fast.
2. Ability to fast video stream ( no live stream simple video stream ).
3. Ability to change quality.

### Non-Functional Requirements
1. Low infrastructure cost.
2. Highly available and scalable.

### Capacity Estimation

- Assume
- 5 million DAU
- DAU watches 5 video upload them.
- 10% of them updated 1 video
- Assume average size of video upload is 300 MB
- So per second -> 50 DAU/second
- Storage per second -> 50 * 0.1 * 300 -> 1.5 GB / second or 150 TB / day.
- etc.

### High level design

So there are two major flow:

1. Streaming video: This can be done by cache. [CDN`] for some pages. If they still need then will have to fetch from our object storage
2. Other apis. this will go to api servers.
3. Video uploading: for this again we will contact api servers but including object storage.

![alt_text](./images/img_1.png)

We are more interested in video stream and uploading part. 

#### Video uploading:

- So video uploading is like putting mp4 or video file on our service. But how our database hasn't been made for this.
- Here's comes the object storage like S3 ( Design this is a separate problem , but were we are using this.).
- So what we can do , from frontend itself directly start uploading the video on S3 ( it would also be better if it do in chunks, to make storage efficient, we can maintain state of chunk updated so retry to previous part only. Basically storing chunks of video ). Once its done we will internaly send the information about the s3 url or details of the video to the api servers that will save these information in db.

![alt_text](./images/img_2.png)

```

# multipart means -> chunk -> paralle

import boto3
from pathlib import Path

s3 = boto3.client('s3')
bucket = "your-bucket-name"
key = "uploads/raw_videos/video.mp4"
file_path = "video.mp4"
chunk_size = 5 * 1024 * 1024  # 5 MB minimum per part

# 1. Initiate multipart upload
response = s3.create_multipart_upload(Bucket=bucket, Key=key)
upload_id = response['UploadId']
parts = []

try:
    with open(file_path, 'rb') as f:
        part_number = 1
        while chunk := f.read(chunk_size):
            response = s3.upload_part(
                Bucket=bucket,
                Key=key,
                PartNumber=part_number,
                UploadId=upload_id,
                Body=chunk
            )
            parts.append({
                'PartNumber': part_number,
                'ETag': response['ETag']
            })
            part_number += 1

    # 2. Complete upload
    s3.complete_multipart_upload(
        Bucket=bucket,
        Key=key,
        UploadId=upload_id,
        MultipartUpload={'Parts': parts}
    )
    print("Upload completed successfully.")

except Exception as e:
    # 3. Abort on failure
    s3.abort_multipart_upload(Bucket=bucket, Key=key, UploadId=upload_id)
    print("Upload failed and aborted:", e)

```

```
s3://your-bucket/uploads/raw_chunks/video_12345/
├── chunk_00001
├── chunk_00002
├── ...
```

![](./images/Screenshot%202025-07-09%20at%2012.38.20%20PM.png)
![](./images/Screenshot%202025-07-09%20at%2012.38.38%20PM.png)

#### Video streaming 

- Fetching data from s3 everytime would be much efficient , better is to keep CDN for this and that will internally connect with s3. [ Note: CDN also incur cost. ]. Streaming protocol e.g. MPEG-DASH , Adobe HTTPS dynamic stream etc. ( HDS ) this is not used much, another HLS. So DASH/HLS . So basically as per these protocol format we send data to s3 and CDN and then it take care. client with these protocal fetch data and play.
- Fetch complete video in one go would be inefficient , its better to break into the chunk of videos. 

![alt_text](./images/img_3.png)

- More about protocol
- In Doubt section
This is how on high level video uploading and stream look like

### Deep-dive high level design (  remeber transcoder also have bitrate so its multplication of 3 variation)

- How cdn getting updated ?
- How we are getting chunks of video ?  
- do we any way to save data while saving video. ? 
- What about different video quality ? 

So more or less we want to little bit deep dive on these part. i.e. Much of video processing.

let's zoom out communication between s3 and CDN.

So Now instead of saving metadata of video via brower. We will save it by backend asynchronously. Since initially we wan't have about video complete information after chunking format etc. Also we don't immediately need this information.

![alt_text](./images/img.png)

1. upload video to s3.
2. that will be moved to transcoding.
3. ``this will chunk the video of each quality`` [ performance can be increase if client sent the chunk of video and object storage save it . So transcoding won't do that] and put in another storage. lets say transcoded storage
4. this transcoded storage will connect to CDN which have video in different quality with chunkings + their thumbnail, watermark etc things.
5. And there would be workers , which will update the meta infomation of video after transcoding completion.
6. We can improve performance by making each thing asynchronously , this will decouple the micro system and can scale them independently. Below if pic that have increase more parallism by distributing responsibility to individuals.

![alt_text](./images/img_4.png)


Video Transcoding itself a complex system. won't discuss it much . We can take it up separately.

![alt_text](./images/img_6.png)

Video Transcoding Responsibilities:

1. Reduce size of video.
2. Have different quality format of chunks of videos.
3. Merging other information like thumbnail etc.

- It is divided into these parts.

![alt_text](./images/img_5.png)

1. Pre-processor
   1. Storge the temporary segmented chunks.
   2. Chunks the video ( if client is not doing ).
   2. Generate DAG config.
2. DAG ( Airflow )
   1. Put the task of different part of video after breaking in video , audio etc. in task queue.
      ![alt_text](./images/img_7.png)
3. Resource manager
   1. It manages all the resource efficiently and help to allocate resource and trigger the worker according to the tasks.
   2. It contain the three queue and task scheduler.
   ![alt_text](./images/img_8.png)
   
4. Workers
   1. Executing the actual task as shown above.
   2. This is also connected to the temporary storage that helps in dealing at particular chunk and merge all task.
5. Encoded video:
   1. It is the final out put . video_chunk_1_{codec}-{resolution}-{bitrates}.mp4

https://github.com/manvirag/tech_learning_tinkering/tree/main/hld/concepts_backup/video_processing

![image](./images/transcoder.png)

## Correction: Transcoding also invole bitrates
- so lets say we decided container mp4 , then it will generate chunks -> resolution * codec * bitrate

- bitrate (resolution is like matrix size, bitrate is like quality in each pixel or box of matrix, mot bits for at box more clarity. )

![](/images/Screenshot%202025-07-09%20at%2011.38.52%20AM.png)
![](/images/Screenshot%202025-07-09%20at%2011.39.00%20AM.png)



![](./images/bitrate.png)
#### References:
1. Alex xu volume 1

#### Doubts:

- what's the protocol used for uploading video ? since it will be long process and how do we do chunking ? First flow from ui to s3.
   - https, we create the chunk of file on frontend and upload on s3 with differen paths , and these according these path we fetch at the time of stream.
- Resolution of number of pixels, more resolution more clarity. What is codec
   - a codec (short for coder-decoder or compressor-decompressor) is a technology or software that compresses and decompresses digital video files. It's essential for reducing file size and making video easier to store, stream, or transmit.

![alt_text](./images/a.png)
![alt_text](./images/b.png)


3. How are we getting the chunks of video from CDN ?
- so eventuall after trancoder -> it will save as per protocol -> basically fragmented mp4, basically small parts of mp4 which are require for streaming in hls. ( .m4s or .ts) , then there is manifest .m3u8 for hls, .mpd  for  Dynamic Adaptive Streaming over HTTP (DASH), dash also require init.mp4

```
s3://your-bucket/videos/abc123/h264/720p/2.5Mbps/chunk_0001.m4s
```

- we use HLS/DASH protocol and fetch the next segment with manifest and run it.
- For e.g. below is list of request for chunk. 
![alt_text](./images/img_9.png)


3. protocol for video upload and video stream ?
- upload: simply https, steam: already mentioned hds,HTTP Live Streaming (HLS) or Dynamic Adaptive Streaming over HTTP (DASH) etc.

4. Are live stream and youtube video are same flow ?
- Answer is yes and no. in case of youtube video its kind of same but here we have complete video, and we can switch to any part of video, its ok to have bit latency and all. in live we do all at time of stream, not previous video ( usually have few minute ), latency matters. though both eventually us HLS and DASH, live stream also use rtmp for ingest without loss. 


Follow up:
- Add live-streaming
- Doubts
  
 
