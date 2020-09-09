use std::char;

fn main() {
    let input = String::from("10#11#12");
    // let input = String::from("1326#");
    // let input = String::from("25#");
    // let input = String::from("12345678910#11#12#13#14#15#16#17#18#19#20#21#22#23#24#25#26#");

    println!("RESULT {}", Solution::freq_alphabets(input));

}

struct Solution;

impl Solution {
    pub fn freq_alphabets(s: String) -> String {
        let mut result = String::new();
        let mut i = s.len() - 1;

        while i >= 0 {
            if s.chars().nth(i).unwrap() == '#' {
                let int = s[i-2..i].parse::<u32>().unwrap();
                result = format!("{}{}", Self::int_to_char(int), result);
                if i <= 2 {
                    break;
                }
                i -= 3;
            } else {
                let int = s.chars().nth(i).unwrap().to_digit(10).unwrap();
                result = format!("{}{}", Self::int_to_char(int), result);
                if i == 0 {
                    break;
                }
                i -= 1
            }

        }

        return result;
    }

    fn int_to_char(i: u32) -> char {
        let x = i + 96;
        return unsafe { char::from_u32_unchecked(x) };
    }
}