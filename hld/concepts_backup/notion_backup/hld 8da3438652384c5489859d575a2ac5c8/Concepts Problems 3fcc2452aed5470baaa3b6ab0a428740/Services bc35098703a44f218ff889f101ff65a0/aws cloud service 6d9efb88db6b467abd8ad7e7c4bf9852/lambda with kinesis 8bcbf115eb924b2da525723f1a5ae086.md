# lambda with kinesis

When ever we connect lambda with another aws service mostly messaging queue.there is a event source mapping whichtake careallthe thinga like fetch from kinesis how and send to lamda we just provide the config fromtop like batch sizeetc. 

this automatically maintain the order of record on basis of partition key . until unless kinessi have wrong order (used putrecords and got partially fail)

Suppose we are connecting kinesis with our service then we will have to handle all the things like fetching handling the order (fetch from single shard not fetch from that particular partition key and dont allowto pro ess until previous record ofsame partitionkey isnt processed. )

kinesis handle the part after in rease/ decrease shard count , only when we are processing with its library