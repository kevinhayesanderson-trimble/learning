fn next_language<'a>(languages: &'a [String], current: &str) -> &'a str {
    let mut found = false;
    for lang in languages {
        if found {
            return lang;
        }
        if lang == current {
            found = true;
        }
    }
    languages.last().unwrap()
}

fn last_language(languages: &[String]) -> &str {
    languages.last().unwrap()
}

fn longest_language<'a>(first: &'a str, second: &'a str) -> &'a str{
    if first.len() >= second.len(){
        first
    }else {
        second
    }
} 

fn main() {
    let languages = vec![
        String::from("rust"),
        String::from("go"),
        String::from("typescript"),
    ];
    let mut result = next_language(&languages, "go");
    print!("{}", result);

    result = last_language(&languages);
    print!("{}", result);

    result = longest_language("go", "typescript");
    print!("{}", result);
}

// struct Account{
//     balance: i32
// }

// struct Bank<'a>{
//     primary_account: &'a Account
// }

// fn make_bank() -> Bank {
//     let account = Account{balance:10};
//     let bank = Bank{primary_account:&account};
//     bank
// }
