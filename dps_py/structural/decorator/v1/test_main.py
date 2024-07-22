import unittest
from main import VowelsRemoveDataSourceDecorator

class MainTest(unittest.TestCase):
    def test_remove_vowels(self):
        cases = [
            ("addition", "ddtn"),
            ("vowels", "vwls"),
            ("self", "slf"),
            #...
        ]
        for one in cases:
            got = VowelsRemoveDataSourceDecorator.remove_vowels(one[0])
            self.assertEqual(got, one[1])
        
        
if __name__ == '__main__':
    unittest.main()