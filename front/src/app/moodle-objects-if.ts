export interface MoodleToken {
    Token: string;
    PrivateToken: string
}

export interface Icon {
    Key: string;
    Component: string;
    Alttext: string;
}

export interface Course {
    Id: number;
    FullName: string;
    ShortName: string;
    IdNumber: string;
    Summary: string;
    SummaryFormat: number;
    StartDate: number;
    EndDate: number;
    Visible: boolean;
    FullNameDisplay: string;
    ViewUrl: string;
    CourseImage: string;
    Progress: number;
    HasProgress: boolean;
    IsFavourite: boolean;
    Hidden: boolean;
    ShowShortName: boolean;
    CourseCategory: string;
}

export interface Subscription {
    DisplayEventsource: boolean;
    SubscriptionName: string;
    SubscriptionUrl: string;
}

export interface Event {
    Id: number;
    Name: string;
    Description: string;
    DescriptionFormat: number;
    Location: string;
    CategoryId?: any;
    GroupId?: any;
    UserId: number;
    RepeatId?: any;
    EventCount?: any;
    ModuleName: string;
    Instance?: any;
    EventType: string;
    TimeStart: number;
    TimeDuration: number;
    TimeSort: number;
    Visible: number;
    TimeModified: number;
    Icon: Icon;
    Course: Course;
    Subscription: Subscription;
    CanEdit: boolean;
    CanDelete: boolean;
    FormattedTime: string;
    IsActionEvent: boolean;
    IsCourseEvent: boolean;
    IsCategoryEvent: boolean;
    GroupName?: any;
    NormalisedEventType: string;
    NormalisedEventTypeText: string;
    Url: string;
    IsLastDay: boolean;
    PopupName: string;
    Draggable: boolean;
}

export interface Date {
    Seconds: number;
    Minutes: number;
    Hours: number;
    mDay: number;
    wDay: number;
    mon: number;
    year: number;
    yDay: number;
    weekday: string;
    month: string;
    timestamp: number;
}

export interface CalenderRequestData {
    Events: Event[];
    DefaultEventContext: number;
    filter_selector: string;
    CourseId: number;
    CategoryId: number;
    IsLoggedIn: boolean;
    Date: Date;
}
