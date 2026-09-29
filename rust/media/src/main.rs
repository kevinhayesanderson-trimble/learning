mod content;
use crate::content::catalog::MightHaveAValue;
use content::media::Media;
use content::media::print_media;
use content::catalog::Catalog;
use content::account::Account;

fn main() {
    let audiobook = Media::Audiobook {
        title: String::from("An Audiobook"),
    };
    print_media(&audiobook);

    let good_movie = Media::Movie {
        title: String::from("Good Movie"),
        director: String::from("Good Director"),
    };
    print_media(&good_movie);

    let bad_book = Media::Book {
        title: String::from("Bad Book"),
        author: String::from("Bad Author"),
    };
    print_media(&bad_book);

    let podcast = Media::Podcast(1);
    print_media(&podcast);

    let placeholder = Media::Placeholder;
    print_media(&placeholder);

    println!("{}", audiobook.description());
    println!("{}", good_movie.description());
    println!("{}", bad_book.description());
    println!("{}", podcast.description());
    println!("{}", placeholder.description());

    let mut catalog = Catalog::new();
    catalog.add(audiobook);
    catalog.add(good_movie);
    catalog.add(bad_book);
    catalog.add(podcast);
    catalog.add(placeholder);

    println!("{:#?}", catalog);
    println!("{:#?}", catalog.items.get(100));
    match catalog.get_by_index_option(0) {
        Some(media) => println!("{:#?}", media),
        None => println!("Nothing at that index"),
    }
    match catalog.get_by_index(0) {
        MightHaveAValue::ThereIsAValue(media) => println!("{:#?}", media),
        MightHaveAValue::NoValueAvailable => println!("Nothing at that index"),
    }

    let item = catalog.get_by_index_option(0);
    //println!("{:#?}", item.unwrap());
    //println!("{:#?}", item.expect("expected to be a item here!"));
    let placeholder = Media::Placeholder;
    println!("{:#?}", item.unwrap_or(&placeholder));

    let mut accounts: Vec<Account> = vec![Account { balance: 0 }, Account { balance: 10 }];

    // field destructuring
    match accounts.first_mut() {
        Some(Account { balance }) => *balance = 30,
        None => println!("No account found"),
    }

    //reference access
    match accounts.first_mut() {
        Some(account) =>{
            account.balance = 300;
        }
        None => println!("No Account Found")
    }
}
