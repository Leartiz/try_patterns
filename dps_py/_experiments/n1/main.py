iis = [
    0.913, 1.202,
    1.082, 0.803
    ]

def trend(t) -> float:
    return 90.59 - 2.773 * t

# ***

def year_results(year) -> list:
    s_count = 4
    ss = year * s_count # 36

    results = []
    for i in range(s_count):
        results.append(
            trend(ss - (s_count - i - 1)) * iis[i])

    return results

# ***

results = year_results(9)

sum = 0
for one in results:
    sum += one

print(sum)


