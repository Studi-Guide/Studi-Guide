export class FeedItem {
    description: string;
    link: string;
    title: string;

    constructor(description: string, link: string, title: string) {
        this.description = description;
        this.link = link;
        this.title = title;
    }
}