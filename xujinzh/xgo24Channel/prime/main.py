import time


def prime(n):
    for num in range(1, n + 1):
        # print(num)
        flag = True
        for i in range(2, num):
            # print("i=", i)
            if num % i == 0:
                flag = False
                break
        if flag == True:
            # print(num)
            pass


if __name__ == "__main__":
    start = time.time()
    prime(800000)
    end = time.time()
    print(f'consume = {end - start} seconds')