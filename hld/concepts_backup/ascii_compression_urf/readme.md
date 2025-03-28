# Encoding (How Data is Represented in Storage & Transfer)

Encoding is the process of converting data into a specific format for efficient storage, transmission, or processing. Different encoding formats offer trade-offs between size, compatibility, and performance.

## **Common Encoding Formats**

| Encoding   | Description                               | Storage Cost & Optimization                         | Example |
|------------|-------------------------------------------|-----------------------------------------------------|---------|
| **ASCII**  | 7-bit character encoding (English only) | ✅ Very compact, ❌ Limited language support        | `'A' -> 01000001` (7-bit) |
| **UTF-8**  | Variable-length encoding (1-4 bytes)    | ✅ Optimized for storage (1 byte for ASCII, more for Unicode) | `'A' -> 01000001` (1 byte), `'€' -> 11100010 10000010 10101100` (3 bytes) |
| **UTF-16** | Fixed 2/4 bytes per character           | ❌ More storage overhead (for ASCII-heavy data)    | `'A' -> 00000000 01000001` (2 bytes), `'€' -> 00100000 10101100` (2 bytes) |
| **Base64** | Converts binary to ASCII (e.g., images in JSON) | ❌ Increases size by ~33%, ✅ Useful for text-based transmission | `'Man' -> TWFu` |
| **Binary Encoding** | Raw data format (e.g., Protobuf, Avro) | ✅ Smallest storage size, ✅ Fast parsing | `Protobuf: message { name: "John" } -> Binary Stream` |

## **What is Binary Data?**
Binary data is **any data represented in raw 0s and 1s** (bits), which is how computers store and process all information. Unlike text formats (e.g., ASCII, JSON), binary data is **not human-readable** but is highly efficient for storage and computation.

### **Types of Binary Data:**
1. **Text in Binary Form** → Even text (e.g., "A") is stored as a binary sequence (`01000001` in ASCII).
2. **Images, Videos, Audio** → JPEG, MP4, MP3 files store pixel, sound, and frame data as raw bytes.
3. **Executable Files & Programs** → Compiled code (e.g., `.exe`, `.bin`) is machine-readable binary.
4. **Structured Data in Binary Formats** → Protocol Buffers (Protobuf), Avro, and MessagePack store structured data in compact binary form.

### **Why Use Binary Data?**
✅ **Smaller storage footprint** → No extra characters like spaces, brackets (unlike JSON, XML).  
✅ **Faster processing** → Directly understood by CPUs, no need for parsing.  
✅ **More efficient transmission** → Takes less bandwidth than text formats.  

## **Base64 Encoding (Text-Based Encoding for Binary Data)**  
Base64 is an encoding scheme that converts **binary data into a text format** using **64 ASCII characters (A-Z, a-z, 0-9, +, /)**. It's commonly used for embedding images in JSON, sending binary files over text-based protocols (e.g., email, HTTP), and avoiding special character issues in transmission.

### **How Base64 Works?**
1. Takes **every 3 bytes (24 bits)** of binary data.  
2. Splits into **four 6-bit groups**.  
3. Maps each group to a **Base64 character**.  
4. If the input isn’t a multiple of 3 bytes, **adds padding (`=`)**.  

#### **Example**  
Binary (`Man` in ASCII):
```
'M' = 01001101
'a' = 01100001
'n' = 01101110
```
Combine: `01001101 01100001 01101110` (24 bits) → Break into 6-bit chunks:
```
010011 010110 000101 101110
```
Map to Base64:
```
'T' 'W' 'F' 'u'  → **TWFu**
```

### **Storage Impact**
- **Base64 increases data size by ~33%** (every 3 bytes become 4 characters).  
- **Useful when transmitting binary in text-based formats** (e.g., images in JSON, email attachments).  

### **When to Avoid Base64?**
- When storage or bandwidth is a concern (e.g., avoid storing images as Base64 in databases, as raw binary is smaller).  

## **Binary Encoding (Compact & Efficient Format for Storage & Transfer)**  
Binary encoding stores **data as raw bytes**, which makes it **more space-efficient** and faster to parse than text-based formats like JSON or XML.

### **How Binary Encoding Works?**
1. **Identify Data Fields:** Convert structured data (e.g., JSON, text, numbers) into a defined schema.  
2. **Use a Compact Representation:** Encode numbers, strings, and objects using the most space-efficient format.  
3. **Store as Raw Bytes:** Instead of using human-readable characters, it directly represents data in a **byte stream**.  
4. **Read with a Decoder:** The receiving system uses a predefined format (e.g., Protobuf schema) to decode the byte stream into meaningful data.  

### **Example: JSON vs. Protobuf (Binary Encoding)**  
#### JSON (Text-Based, More Space)  
```json
{
  "name": "John",
  "age": 25
}
```
Storage: ~26 bytes (including formatting characters, field names, etc.)  

#### Protobuf (Binary Format, More Efficient)  
```proto
message User {
  string name = 1;
  int32 age = 2;
}
```
Serialized Binary Data:
```
08 4A 10 19 (Hex representation)
```
Storage: **Much smaller (~8 bytes)** due to **efficient field encoding**.  

### **Advantages of Binary Encoding**
✅ **Smaller storage footprint** → Reduces space usage significantly.  
✅ **Faster parsing & transmission** → No need to process human-readable text.  
✅ **Ideal for structured data** → Used in Avro, Protobuf, MessagePack, and gRPC.  

### **When to Use Binary Encoding?**
✅ When **storing large structured data** (e.g., user profiles, logs, analytics).  
✅ When **transferring data between services** (e.g., gRPC uses Protobuf).  
✅ When **speed & efficiency matter** (e.g., replacing JSON for internal APIs).  

### **When NOT to Use Binary Encoding?**
❌ When **human readability is required** (e.g., configuration files, APIs meant for debugging).  
❌ When **interoperability with text-based systems is needed** (e.g., browsers don’t natively support Protobuf).  

## **Key Takeaways**
✅ **UTF-8** is the most efficient for text storage and transfer.  
✅ **Base64 is useful for text-based transmission of binary data** (but adds 33% overhead).  
✅ **Binary encoding is best for compact, efficient storage & transmission** (e.g., Protobuf, Avro).  
✅ **Use Base64 only when necessary** (e.g., embedding images in JSON, email attachments).  
✅ **For structured data, prefer Protobuf/Avro over JSON** to save space and speed up processing.  

Encoding impacts storage costs and performance, so choosing the right format is essential for an optimized system! 🚀


