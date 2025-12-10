import sys
import pulp as pl

def parse_input():
    indicators = []
    buttons = []
    joltage = []

    for line in sys.stdin:
        parts = line.strip().split()

        bitstring = parts[0][1:-1].replace('.', '0').replace('#', '1')
        indicators.append((int(bitstring, 2), len(parts[0]) - 2))

        group = []
        for token in parts[1:-1]:
            group.append(list(map(int, token[1:-1].split(','))))
        buttons.append(group)

        joltage.append(tuple(map(int, parts[-1][1:-1].split(','))))

    return indicators, buttons, joltage


def solve_instance(buttons_i, target):
    prob = pl.LpProblem("min_sum", pl.LpMinimize)
    x = [pl.LpVariable(f"x_{i}", lowBound=0, cat="Integer") for i in range(len(buttons_i))]

    for d, t in enumerate(target):
        prob += sum((d in btn) * x[j] for j, btn in enumerate(buttons_i)) == t

    prob += sum(x)
    prob.solve(pl.PULP_CBC_CMD(msg=0))

    return sum(v.value() for v in x)


def main():
    indicators, buttons, joltage = parse_input()
    total = 0
    for b, j in zip(buttons, joltage):
        total += solve_instance(b, j)
    print(total)


if __name__ == "__main__":
    main()
