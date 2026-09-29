use std::vec;


fn print_elements(elements: &[String]){
    // for element in elements{
    //     println!("{:#?}", element);
    //     }

    elements
    .iter()
    .map(|el| format!("{} {}", el, el)) //iterator adaptor
    .for_each(|el| println!("{}", el)); //iterator consumer
}

fn shorten_strings(elements: &mut[String]) {
    elements
        .iter_mut()
        .for_each(|el| el.truncate(1));
}

fn to_uppercase(elements: &[String]) -> Vec<String>{
    elements
        .iter()
        .map(|el| el.to_uppercase())
        .collect::<Vec<String>>()
}

fn move_elements(vec_a: Vec<String>, vec_b: &mut Vec<String>){
    vec_a
        .into_iter()
        .for_each(|el| vec_b.push(el));
}

fn explode(elements: &[String]) -> Vec<Vec<String>>{
    elements
        .iter()
        .map(|el| el.chars().map(|c| c.to_string()).collect())
        .collect()
}
fn find_color_or(elements: &[String], search: &str, fallback: &str) -> String{
    elements
        .iter()
        .find(|el| el.contains(search))
        .map_or(String::from(fallback), |el | el.to_string())
}
fn main() {
    let mut colors = vec![
        String::from("red"),
        String::from("green"),
        String::from("blue")
    ];
    print_elements(&colors);
    print_elements(&colors[1..3]);
    shorten_strings(&mut colors);
    println!("{:#?}", colors);
    // let mut colors_iter = colors.iter();
    // println!("{:#?}", colors_iter.next());
    // println!("{:#?}", colors_iter.next());
    // println!("{:#?}", colors_iter.next());
    // println!("{:#?}", colors_iter.next());

    let uppercased = to_uppercase(&colors);
    println!("{:#?}", uppercased);

    let mut destination =  vec![];
    move_elements(colors, &mut destination);
    println!("{:#?}", destination);

    let colors1 = vec![
        String::from("red"),
        String::from("green"),
        String::from("blue")
    ];
    let exploded = explode(&colors1);
    println!("{:#?}", exploded);

    let mut found_color = find_color_or(&colors1, "purple", "orange");
    println!("{:#?}", found_color);
    found_color = find_color_or(&colors1, "bl", "orange");
    println!("{:#?}", found_color);
}