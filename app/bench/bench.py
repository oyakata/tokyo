# -*- coding:utf-8 -*-
import random


def countup1(items, targets):
    L = []
    for target in targets:
        count = 0
        for x in items:
            if x == target:
                count += 1
        L.append((target, count))
    return L


def countup2(items, targets):
    L = []
    tmp = {x: 0 for x in targets}

    for x in items:
        tmp[x] += 1
    return [(x, tmp[x]) for x in targets]


def make_params(x, y):
    return (
        [random.randint(1, y) for i in range(x)],
        list(range(1, y + 1)),
    )


P1A, P1B = make_params(50000, 50)
P2A, P2B = make_params(50000, 50000)
P3A, P3B = make_params(50, 50000)


if __name__ == '__main__':
    import timeit

    items = (
        ('P1A', 'P1B', (500, 500)),
        ('P2A', 'P2B', (1, 100)),
        ('P3A', 'P3B', (500, 1000)),
    )

    G = globals()

    # 1(50000, 50)>   500 46.40978667300078s
    # 2(50000, 50)>   500 2.749026679999588s
    # 1(50000, 50000)>    1   94.42981744899953s
    # 2(50000, 50000)>    100 2.782400592999693s
    # 1(50, 50000)>   500 51.12394202299947s
    # 2(50, 50000)>   1000    13.445461807001266s

    for x, y, ns in items:
        one, two = ns
        result1 = timeit.timeit(
            'countup1({}, {})'.format(x, y),
            setup='from __main__ import countup1, P1A, P1B, P2A, P2B, P3A, P3B',
            number=one,
        )
        print('1({}, {})>\t{}\t{}s'.format(len(G[x]), len(G[y]), one, result1))

        result2 = timeit.timeit(
            'countup2({}, {})'.format(x, y),
            setup='from __main__ import countup2, P1A, P1B, P2A, P2B, P3A, P3B',
            number=two,
        )
        print('2({}, {})>\t{}\t{}s'.format(len(G[x]), len(G[y]), two, result2))
