# Benchmark Serialization Methods - GoLang

## Environment

**SO:** macOS Big Sur 11.1

**Processor:** 2,6 GHz 6-Core Intel Core i7

**Memory:** 16 GB 2400 MHz DDR4

**GoLang:** v1.14

**Protobuf Lib:** v1.4.1


## Data

```
{
    "name":"John Doe",
    "id":1234,
    "email":"jdoe@example.com",
    "phones":[
        {
            "number":"555-4321",
            "type":1
        }
    ]
}
```

## Size

| **Protobuf**  | **JSON** | **XML** |
| ------------- |:----:| -----:|
| 45 bytes | 98 bytes | 138 bytes |


## Time Spent

### Serializing

| **Protobuf**  | **JSON** | **XML** |
| ------------- |:----:| -----:|
| 207µs | 200µs | 100µs |
| 291µs | 262µs | 87µs |
| 222µs | 317µs | 60µs |
| 257µs | 189µs | 74µs |
| 282µs | 341µs | 73µs |
| 190µs | 137µs | 50µs |
| 178µs | 201µs | 70µs |
| 207µs | 143µs | 52µs |
| 257µs | 149µs | 61µs |
| 262µs | 192µs | 66µs |
| ---- | ---- | ---- |
| **Avg 235µs** | **Avg 213µs** | **Avg 70µs** |

### Deserializing

| **Protobuf**  | **JSON** | **XML** |
| ------------- |:----:| -----:|
| 194µs |  141µs | 57µs |
| 202µs |  152µs | 60µs |
| 223µs |  214µs | 69µs |
| 280µs |  173µs | 94µs |
| 263µs |  229µs | 80µs |
| 281µs |  156µs | 63µs |
| 248µs |  200µs | 64µs |
| 196µs |  147µs | 70µs |
| 204µs |  171µs | 60µs |
| 207µs |  204µs | 107µs |
| ---- | ---- | ---- |
| **Avg 230µs** | **Avg 178µs** | **Avg 72µs** |

## Conclusions

- `XML` is about 70% faster than `Protobuf`, in the serialization process.
- `Protobuf` reduces the size of the information structure by 40% to 70%.
- Both process has a similar time consumption.
- CPU usage was not measured during the exercise 
- The code complexity is same for different methods using GoLang.