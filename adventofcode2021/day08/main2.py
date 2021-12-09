import dataclasses
import fileinput
from typing import ClassVar


#  0:      1:      2:      3:      4:
#  aaaa    ....    aaaa    aaaa    ....
# b    c  .    c  .    c  .    c  b    c
# b    c  .    c  .    c  .    c  b    c
#  ....    ....    dddd    dddd    dddd
# e    f  .    f  e    .  .    f  .    f
# e    f  .    f  e    .  .    f  .    f
#  gggg    ....    gggg    gggg    ....
#
# 5:      6:      7:      8:      9:
#  aaaa    aaaa    aaaa    aaaa    aaaa
# b    .  b    .  .    c  b    c  b    c
# b    .  b    .  .    c  b    c  b    c
#  dddd    dddd    ....    dddd    dddd
# .    f  e    f  .    f  e    f  .    f
# .    f  e    f  .    f  e    f  .    f
#  gggg    gggg    ....    gggg    gggg


@dataclasses.dataclass
class Digits:
    # digits: dict[int, str] = dataclasses.field(default_factory=dict)
    alphabet: ClassVar[str] = "abcdefg"
    mapping: ClassVar[dict[int, list[int]]] = {
        0: [0, 1, 2, 4, 5, 6],
        1: [2, 5],
        2: [0, 2, 3, 4, 6],
        3: [0, 2, 3, 5, 6],
        4: [1, 2, 3, 5],
        5: [0, 1, 3, 5, 6],
        6: [0, 1, 3, 4, 5, 6],
        7: [0, 2, 5],
        8: [0, 1, 2, 3, 4, 5, 6],
        9: [0, 1, 2, 3, 5, 6],
    }
    segments: list[str] = dataclasses.field(default_factory=lambda: [""] * 7)

    def __setitem__(self, digit: int, segments: str) -> None:
        if digit not in range(9):
            raise IndexError()
        # if digit in self.digits:
        #     raise Exception(f"Digit {digit} is already defined")
        self.check_segments(segments)

        mapping = Digits.mapping[digit]
        if len(mapping) != len(segments):
            raise Exception("Unexpected length")

        # self.digits[digit] = segments
        for i, m in enumerate(mapping):
            self.segments[m] = segments[i]  # todo check if already assigned and not equal

    def __getitem__(self, item: int) -> str:
        return ''.join([self.segments[i] for i in self.mapping[item]])

    def count_empty(self):
        return sum([1 for i in self.segments if i == ''])


    def digit(self, segments) -> [int | None]:
        print(segments)
        print(self.segments)
        indices = [self.segments.index(s) for s in segments]
        for k, v in self.mapping:
            if v == indices:
                return k
        return None

    def display(self, digit) -> str:
        return """
 {0}{0}{0}{0} 
{1}    {2}
{1}    {2}
 {3}{3}{3}{3} 
{4}    {5}
{4}    {5}
 {6}{6}{6}{6} 
 """.format(*[s if (s != '' and i in self.mapping[digit]) else '.' for i, s in enumerate(self.segments)])

    @staticmethod
    def check_segments(segments):
        for segment in segments:
            if len(segment) != 1 and segment not in Digits.alphabet:
                raise Exception(f"Unexpected segment {segment}")


def run(seq) -> int:
    count = 0
    for s in seq:
        guess(s.split())
    return count


def guess(data) -> Digits:
    d = Digits()
    for segment in data[:-5]:
        if len(segment) == 2:
            d[1] = segment
        if len(segment) == 3:
            d[7] = segment
        if len(segment) == 4:
            d[4] = segment
        if len(segment) == 7:
            d[8] = segment

    print()
    print(d.count_empty())

    for segments in data[-4:]:
        print(d.digit(segments))
        print(segments)
    return d


if __name__ == '__main__':
    print(run(fileinput.input()))


def test_get_set_items_1():
    d = Digits()
    d[1] = 'be'
    assert d[1] == 'be'
    assert d.segments == ['', '', 'b', '', '', 'e', '']


def test_get_set_items_2():
    d = Digits()
    d[8] = 'abcdefg'

    assert d[8] == 'abcdefg'
    assert d.segments == ['a', 'b', 'c', 'd', 'e', 'f', 'g']


def test_guess():
    d = guess("be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe".split())
    print(d.count_empty())
    # print(d)


testdata = """be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe
edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc
fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg
fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb
aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea
fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb
dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe
bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef
egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb
gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce
"""


def test_data():
    result = run(testdata.splitlines())
    assert result == 61229
