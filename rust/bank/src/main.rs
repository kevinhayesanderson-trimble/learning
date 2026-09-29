#[derive(Debug)]
struct Account {
    id: u32,
    balance: i32,
    holder: String,
}

impl Account {
    fn new(id: u32, holder: String) -> Self {
        Account {
            id,
            holder,
            balance: 0,
        }
    }

    fn deposit(&mut self, amount: i32) -> i32 {
        self.balance += amount;
        self.balance
    }

    fn withdraw(&mut self, amount: i32) -> i32 {
        self.balance -= amount;
        self.balance
    }

    fn summary(&self) -> String {
        format!(
            "{}-{} has a balance of {}",
            self.holder, self.id, self.balance
        )
    }
}

#[derive(Debug)]
struct Bank {
    accounts: Vec<Account>,
}

impl Bank {
    fn new() -> Self {
        Bank { accounts: vec![] }
    }

    fn add_account(&mut self, account: Account) {
        self.accounts.push(account);
    }

    fn total_balance(&self) -> i32 {
        self.accounts
            .iter()
            .map(|account| account.balance as i32)
            .sum()
    }

    fn summary(&self) -> Vec<String> {
        self.accounts
            .iter()
            .map(|account| {
                format!(
                    "{}-{} has a balance of {}",
                    account.holder, account.id, account.balance
                )
            })
            .collect::<Vec<String>>()
    }
}

// fn print_account(account: &Account) {
//     println!("{:#?}", account);
// }

// fn print_holder(holder: String) {
//     println!("{}", holder);
// }

// fn change_account(account: &mut Account) {
//     account.balance = 10;
// }

// fn add_account(bank: &mut Bank, account: Account) {
//     bank.accounts.push(account);
// }

// fn make_and_print_account() -> &Account {
//     let account = Account::new(2, String::from("me1"));
//     println!("{:#?}", account);
//     &account
// }
fn main() {
    let mut bank = Bank::new();
    let mut account = Account::new(1, String::from("me"));

    account.deposit(500);
    account.withdraw(250);
    println!("{}", account.summary());
    //print_holder(account.holder);

    bank.add_account(account);
    println!("{:#?}", bank.total_balance());
    println!("{:#?}", bank.summary());
    //let account_ref_1 = &account;
    //let account_ref_2 = account_ref_1;

    //print_account(&account);
    //print_account(account_ref_1);
    //print_account(account_ref_2);

    //println!("{:#?}", account);

    //let other_account = account;
    //print_account(&other_account);

    //change_account(&mut account);

    //add_account(&mut bank, account);

    //println!("{:#?}", bank);

    // some values are copied instead of moved
    //all numbers, bool, char, Arrays,Tuples and references(both readable and writable)
    //let num = 5;
    //let other_num = num; //5 is copied not moved
    //println!("{} {}", num, other_num);

    // let account_ref = make_and_print_account();
    // println!("{:#?}", account_ref);
}
