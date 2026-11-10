Information:

Quantifying INformation:

Probalility of data to its occurence:
- usually high probability data will have low information content
- low probability data will have the highest information content - thus we assign large length of bit strings while encoding these.
- 

Encoding:
- data to bit strings
- unabigious encoding can be visualised in a binary tree:
  - all leaves have one parent node
  - 

Entropy:
- minimal size of the encoding needed to represent the data.information content
- theoritical minimum
- 

Variable length encoding:
- idea is to assign lower size bits to the most occuring data, and higher size bits to the low occuring data
- must be close to the entropy for efficiency

Huffman's Algorithm:
- method to contruct a optimal variable-length encoding
- optimal binary tree built form the bottom up
- one symbol at a time, each node represent a single symbol
- 

Huffman's Code:
- we can optimise by choosing the pairs of symbol for encoding, reducing the length and packing more data in the encoding
- eg: modern file compression algorithm, identifies the most occuring pattern of data aand encodes it to represent the smallest bits
- https://www.geeksforgeeks.org/computer-networks/lzw-lempel-ziv-welch-compression-technique/

Error detection and correction:
- Simple encoding techniques doesn;t help with the errors
- Hamming Distance:
  - distance by which there is a bit flip
  - Hamming distance power-ups error detection and correction: idea is to increase the hamming distance
  - Single bit error Detection:
      - minimun hamming distance is 2
      - parity bit is added to the match the required length
      - even parity - 1 bit is added and no of 1 bit is even -  this checking the no of one's will tell us about the corruption(parity error)
      - odd parity -
      - minimum hamming distance = code words + parity bits
      - XOR is used to perform parity checks - XOR'ing all the bits in the event parity scheme will yield 0 for non errors
      - Limitations:
          - Parity bits can only detect an odd number of bit errors (1, 3, 5, etc.).
          - If two bits flip at the same time, the count of 1s changes by two, leaving the total even/odd status unchanged. The receiver will be completely blind to the error.
   
  - Multiple bit error Detection
    - for detecting E errors(eg: 2bit flip), minimum hamming distance is E+1 between the words
    - [ Data Block ] ──► Strategy Choice:
                   ├── 1. Two-Dimensional Parity (Grid Check)
                        - row and column cheksums -
                        -  2 and 3-bit errors, missed 4-bit error
                   ├── 2. Checksum (Arithmetic Summing)
                        - binary add data and attach the checksum to the end of the trans mission, 1 - ok, 0 - error
                   └── 3. Cyclic Redundancy Check / CRC (Polynomial Division)
                        - standard
                        - Binary long division - XOR subtraction
                        -
- Error Correction:
  - for correction of E erros, hamming distance = 2E+1 between the words
  - coding theory 
      


