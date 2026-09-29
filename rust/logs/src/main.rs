use std::fs;
use std::io::Error;

fn extract_errors(text: &str) -> Vec<String> {
    let split_text = text.split("\n");

    let mut results = vec![];

    for line in split_text {
        if line.starts_with("ERROR") {
            results.push(line.to_string());
        }
    }

    results
}

fn main() -> Result<(), Error> {

    // //1st way
    //let mut error_logs = vec![];
    // match fs::read_to_string("logs.txt") {
    //     Ok(text) => {
    //         println!("{:#?}", text.len());
    //         error_logs = extract_errors(text.as_str());
    //         match fs::write("errors.txt", error_logs.join("\n")) {
    //             Ok(..) => println!("Wrote errors.txt"),
    //             Err(error_writing_file) => println!("Error while writing errors.txt: {:#?}", error_writing_file)
    //         }
    //     }
    //     Err(error) => println!("Failure in reading file: {:#?}", error),
    // }

    // //2nd way
    // let text = fs::read_to_string("logs.txt").expect("failed to read logs.txt");
    // let error_logs = extract_errors(text.as_str());
    // fs::write("errors.txt", error_logs.join("\n")).expect("failed to write errors.txt");

    // 3rd way
    let text = fs::read_to_string("logs.txt")?;
    let error_logs = extract_errors(text.as_str());
    fs::write("errors.txt", error_logs.join("\n"))?;
    println!("{:#?}", error_logs);
    Ok(())
    //println!("{:#?}", text);
    // match divide(5.0, 3.0) {
    //     Ok(quotient) => {
    //         println!("{}", quotient)
    //     }
    //     Err(what_went_wrong) => {
    //         println!("{}", what_went_wrong)
    //     }
    // }
    // let color = "red";
    // println!("{:#?}", color);
    // println!("{:#?}", &color[1..2]);
    // println!("{:#?}", color);
    // println!("{:#?}", &color[0..2]);
}

fn divide(dividend: f64, divisor: f64) -> Result<f64, Error> {
    if divisor == 0.0 {
        Err(Error::other("can't divide by zero"))
    } else {
        Ok(dividend / divisor)
    }
}

// fn build_color() -> &str {
//     let color = String::from("red");

//     // 'as_str()' creates a string slice that refers to the String's heap-allocated data
//     color.as_str()
// }
