import pandas as pd 
import matplotlib.pyplot as plt 

df = pd.read_csv("algorithm-race.csv")

plt.figure(figsize=(10, 5))
plt.plot(df['SearchArraySize'], df['HtreeTime'], label='HtreeTime', marker='o')
plt.plot(df['SearchArraySize'], df['LSTime1000'], label='LSTime1000', marker='s')

plt.xlabel('Search Array Size')
plt.ylabel('Time (ms)')
plt.title('Benchmarking: Search Array Size vs Elapsed Time')
plt.legend()
plt.grid(True)
plt.show()
