interface Source {
    id: number,
    name: string
    url: string
    searchUrl: string
    header: string
}

interface Search {
    bookName: number
    author: string
    newChapter: string
    bookUrl: string
}

export {
    Source,
    Search
}