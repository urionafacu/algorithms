pub fn is_palindrome(text: &str) -> bool {
	if text.len() == 0 {
		return false;
	}

	let text: Vec<char> = text.chars().collect();

	let mut first_index: usize = 0;

	let mut last_index: usize = text.len() - 1;

	while first_index < last_index {
		if !text[first_index].is_alphabetic() {
			first_index += 1;
			continue;
		}
		if !text[last_index].is_alphabetic() {
			last_index -= 1;
			continue;
		}

		if text[first_index].to_ascii_lowercase() != text[last_index].to_ascii_lowercase() {
			return false;
		}

		first_index += 1;
		last_index -= 1;
	}

	true
}

fn main() {
	let ala: String = String::from("ala");

	let response: bool = is_palindrome(&ala);

	println!("{:?} is palindrome ala?", response);
}