# -- coding: utf-8 --


def count_to(count):
    numbers = ["one", "two", "three", "four", "five"]
    for pos, number in zip(range(count), numbers):
        yield number

count_to_two = lambda: count_to(2)
count_to_five = lambda: count_to(5)

print('Counting to two...')
for number in count_to_two():
    print number


print('Counting to five...')
for number in count_to_five():
    print number
