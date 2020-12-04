export class FeedItem {
    description: string;
    link: string;
    title: string;
    content: string;
    published: Date;
    isExpanded: boolean;
    rootUrl: string;
    constructor(description: string, link: string, rootUrl:string, title: string, published: Date,content:string) {
        this.description = description;
        this.link = link;
        this.title = title;
        this.published = published;
        this.rootUrl = rootUrl;
        this.content = this.cleanUpContent(content);
    }

    private cleanUpContent(content: string) {
        // Cleanup relative hrefs to link to main root page
        return content.replace('href="/', 'href="' + this.rootUrl + '/');
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