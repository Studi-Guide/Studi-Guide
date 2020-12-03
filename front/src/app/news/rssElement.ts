export class FeedItem {
    description: string;
    link: string;
    title: string;
    content: string;
    published: Date;
    constructor(description: string, link: string, title: string, published: Date,content:string) {
        this.description = description;
        this.link = link;
        this.title = title;
        this.published = published;
        this.content = content;
    }
}

export class Feed {
    title: string;
    description: string;
    link: string;
    published: Date;

    constructor(description: string, link: string, title: string, published: Date) {
        this.description = description;
        this.link = link;
        this.title = title;
        this.published = published;
    }
}