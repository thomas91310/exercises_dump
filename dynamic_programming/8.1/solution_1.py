"""count different ways to go walk up the stairs"""

def count(n_steps):
    """count counts the number of possibilities
    given the number of steps the child needs to go
    """
    mem = {}
    return recursion(mem, 0, 0, n_steps, 1) + \
            recursion(mem, 0, 0, n_steps, 2) +  \
            recursion(mem, 0, 0, n_steps, 3)

def recursion(memoize, steps_so_far, n_possibilities, n_steps, to_add):
    """Recursion with memoization"""
    print(memoize)
    steps_so_far += to_add
    if memoize.get(steps_so_far, -1) != -1:
        return memoize[steps_so_far]

    if steps_so_far == n_steps:
        return n_possibilities + 1
    elif steps_so_far > n_steps:
        return n_possibilities

    memoize[steps_so_far] = recursion(memoize, steps_so_far, n_possibilities, n_steps, 1) + \
        recursion(memoize, steps_so_far, n_possibilities, n_steps, 2) + \
        recursion(memoize, steps_so_far, n_possibilities, n_steps, 3)

    return memoize[steps_so_far]

def main():
    """main"""
    print(count(5))

if __name__ == '__main__':
    main()
