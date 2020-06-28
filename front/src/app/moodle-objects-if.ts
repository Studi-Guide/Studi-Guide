export interface MoodleToken {
    token: string;
    privateToken: string
}

export interface Icon {
    key: string;
    component: string;
    alttext: string;
}

export interface Course {
    id: number;
    fullname: string;
    shortname: string;
    idnumber: string;
    summary: string;
    summaryformat: number;
    startdate: number;
    enddate: number;
    visible: boolean;
    fullnamedisplay: string;
    viewurl: string;
    courseimage: string;
    progress: number;
    hasprogress: boolean;
    isfavourite: boolean;
    hidden: boolean;
    showshortname: boolean;
    coursecategory: string;
}

export interface Subscription {
    displayeventsource: boolean;
    subscriptionname: string;
    subscriptionurl: string;
}

export interface Event {
    id: number;
    name: string;
    description: string;
    descriptionformat: number;
    location: string;
    categoryid?: any;
    groupid?: any;
    userid: number;
    repeatid?: any;
    eventcount?: any;
    modulename: string;
    instance?: any;
    eventtype: string;
    timestart: number;
    timeduration: number;
    timesort: number;
    visible: number;
    timemodified: number;
    icon: Icon;
    course: Course;
    subscription: Subscription;
    canedit: boolean;
    candelete: boolean;
    deleteurl: string;
    editurl: string;
    viewurl: string;
    formattedtime: string;
    isactionevent: boolean;
    iscourseevent: boolean;
    iscategoryevent: boolean;
    groupname?: any;
    normalisedeventtype: string;
    normalisedeventtypetext: string;
    url: string;
    islastday: boolean;
    popupname: string;
    draggable: boolean;
}

export interface Date {
    seconds: number;
    minutes: number;
    hours: number;
    mday: number;
    wday: number;
    mon: number;
    year: number;
    yday: number;
    weekday: string;
    month: string;
    timestamp: number;
}

export interface CalenderRequestData {
    events: Event[];
    defaulteventcontext: number;
    filter_selector: string;
    courseid: number;
    categoryid: number;
    isloggedin: boolean;
    date: Date;
}
