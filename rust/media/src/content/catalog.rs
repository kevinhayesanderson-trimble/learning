use super::media::Media;

#[derive(Debug)]
pub struct Catalog {
    pub items: Vec<Media>,
}

impl Catalog {
    pub fn new() -> Self {
        Catalog { items: vec![] }
    }
    pub fn add(&mut self, media: Media) {
        self.items.push(media);
    }
    pub fn get_by_index_option(&self, index: usize) -> Option<&Media> {
        // Before: Redundant unpacking and repacking
        // match self.items.get(index) {
        //     Some(media) => Some(media),
        //     None => None,
        // }
        // After: The idiomatic Rust way
        self.items.get(index)
    }
    pub fn get_by_index(&self, index: usize) -> MightHaveAValue<'_> {
        if self.items.len() > index {
            MightHaveAValue::ThereIsAValue(&self.items[index])
        } else {
            MightHaveAValue::NoValueAvailable
        }
    }
}

pub enum MightHaveAValue<'a> {
    ThereIsAValue(&'a Media),
    NoValueAvailable,
}