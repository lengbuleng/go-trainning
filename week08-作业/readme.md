1、使用 redis benchmark 工具, 测试 10 20 50 100 200 1k 5k 字节 value 大小，redis get set 性能。

1）测试 10 字节value大小
$redis-benchmark -d 10 -t get,set

结果如下：
====== SET ======
  100000 requests completed in 0.73 seconds, 136239.78 requests per second
====== GET ======
  100000 requests completed in 0.80 seconds, 125000.00 requests per second

2）测试 20 字节value大小
$redis-benchmark -d 20 -t get,set

结果如下：
====== SET ======
  100000 requests completed in 0.70 seconds, 143678.17 requests per second
====== GET ======
  100000 requests completed in 0.75 seconds, 132978.73 requests per second

3）测试 50 字节value大小
$redis-benchmark -d 50 -t get,set

结果如下：
====== SET ======
  100000 requests completed in 0.79 seconds, 126103.41 requests per second
====== GET ======
  100000 requests completed in 0.80 seconds, 125156.45 requests per second

4）测试 100 字节value大小
$redis-benchmark -d 100 -t get,set

结果如下：
====== SET ======
  100000 requests completed in 0.71 seconds, 140449.44 requests per second
====== GET ======
  100000 requests completed in 0.81 seconds, 123915.74 requests per second

5）测试 200 字节value大小
$redis-benchmark -d 200 -t get,set

结果如下：
====== SET ======
  100000 requests completed in 0.70 seconds, 142247.52 requests per second
====== GET ======
  100000 requests completed in 0.66 seconds, 151515.14 requests per second

6）测试 1k 字节value大小
$redis-benchmark -d 1024 -t get,set

结果如下：
====== SET ======
  100000 requests completed in 0.75 seconds, 134228.19 requests per second
====== GET ======
  100000 requests completed in 0.70 seconds, 143266.47 requests per second

7）测试 5k 字节value大小
$redis-benchmark -d 5120 -t get,set

结果如下：
====== SET ======
  100000 requests completed in 0.78 seconds, 128205.13 requests per second
====== GET ======
  100000 requests completed in 0.73 seconds, 136986.30 requests per second


2、写入一定量的 kv 数据, 根据数据大小 1w-50w 自己评估, 结合写入前后的 info memory 信息 , 分析上述不同 value 大小下，平均每个 key 的占用内存空间。

1）先清空redis缓存，保证写入前redis缓存占用内存为0
$redis-cli -h 127.0.0.1 -p 6379
$flushall

2）执行./main，查看生成的reports下的分析报告

Key,Count,Size,NeverExpire,AvgTtl(excluded never expire)
len5120_10w:*,100000,488.472 MB,100000,0
len1024_10w:*,100000,97.847 MB,100000,0
len5120_1w:*,10000,48.847 MB,10000,0
len1024_1w:*,10000,9.785 MB,10000,0
len10_10w:*,100000,1.049 MB,100000,0
len10_1w:*,10000,107.422 KB,10000,0

3）查看不同的value大小下，平均每个key的占用内存空间。对整个redis进行扫描，寻找较大的key

$redis-cli -h 127.0.0.1 -p 6379 --bigkeys

结果如下：

 Scanning the entire keyspace to find biggest keys as well as
 average sizes per key type.  You can use -i 0.1 to sleep 0.1 sec
 per 100 SCAN commands (not usually needed).

[00.00%] Biggest string found so far 'len10_10w:86916' with 10 bytes
[00.00%] Biggest string found so far 'len1024_10w:74786' with 1024 bytes
[00.00%] Biggest string found so far 'len5120_10w:79884' with 5120 bytes

-------- summary -------

Sampled 330000 keys in the keyspace!
Total key length in bytes is 5293340 (avg len 16.04)

Biggest string found 'len5120_10w:79884' has 5120 bytes

0 lists with 0 items (00.00% of keys, avg size 0.00)
0 hashs with 0 fields (00.00% of keys, avg size 0.00)
330000 strings with 676940000 bytes (100.00% of keys, avg size 2051.33)
0 streams with 0 entries (00.00% of keys, avg size 0.00)
0 sets with 0 members (00.00% of keys, avg size 0.00)
0 zsets with 0 members (00.00% of keys, avg size 0.00)