"""find_path finds all the paths to get to he bottom right corner
of the grid
"""
import copy

def find_path(current_row, current_col, current_path, rows, columns, \
                    forbidden_cells, right, down, paths):
    """Recursion"""
    forbidden = forbidden_cells.get(current_row, None)
    if forbidden is not None and current_col in forbidden:
        return None

    current_path.append((current_row, current_col))
    print("row : ", current_row)
    print("col : ", current_col)
    print("path : ", current_path)

    if current_row + 1 == rows and current_col + 1 == columns:
        paths.append(current_path)
        return paths

    if current_row + 1 < rows:
        copy_current = copy.deepcopy(current_path)
        find_path(current_row + 1, current_col, copy_current, rows, columns, \
        forbidden_cells, right, down, paths)

    if current_col + 1 < columns:
        copy_current = copy.deepcopy(current_path)
        find_path(current_row, current_col + 1, copy_current, rows, columns, \
         forbidden_cells, right, down, paths)

    return paths

def main():
    """main"""
    rows = 5
    columns = 5
    forbidden_cells = make_forbidden_cells()
    current_path, right_recurs, down_recurs, paths = [], [], [], []
    paths = find_path(0, 0, current_path, rows, columns, forbidden_cells, \
                         paths, right_recurs, down_recurs)
    print(paths)

def make_forbidden_cells():
    """make_forbidden_cells returns a datastructure that represents
    the cells the robot can't step on
    """
    return {
        1: [1, 4],
        2: [8],
        3: [4],
        # 4: [3],
        9: [7],
    }

if __name__ == '__main__':
    main()
