# matching-engine-heap
Python implementation of a matching engine using binary heaps for maintaining exchange orderbook

Execute Test1.py and Test2.py files.

The value of 'order_id' reflects the order of arrival to the exchange. Orders are executed in the priority of the 'order_id'. Orders with a smaller 'order_id', other conditions being equal, get to be executed first, because these are the orders which arrived earlier.

Both of binary heaps, 'buys_prices' and 'sells_prices', are min-heaps. 'buys_prices' heap contains minus of submitted buy prices. The binary heaps here contain unique entries. When we store information about new orders at each price level in the 'buys_info' and 'sells_info' we check first whether that price is already there. We can use that check to also see whether we need to push that price into the corresponding heap, and avoid duplicates.


