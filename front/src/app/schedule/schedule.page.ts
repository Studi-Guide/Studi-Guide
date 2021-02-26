import {AfterViewInit, Component, ViewChild} from '@angular/core';
import {MoodleService} from '../services/moodle.service';
import {Event, MoodleToken} from '../moodle-objects-if';
import {LoadingController} from '@ionic/angular';
import {Router} from '@angular/router';
import {LoginComponent} from './login/login.component';
import {Storage} from '@ionic/storage';


@Component({
  selector: 'app-schedule',
  templateUrl: 'schedule.page.html',
  styleUrls: ['schedule.page.scss']
})
export class SchedulePage implements AfterViewInit {

  constructor(
      private moodleService: MoodleService,
      public loadingController: LoadingController,
      private router: Router,
      // TODO remove when presentation is done
      private storage: Storage
  ) {}
  private token: MoodleToken;
  public calenderEvents: Event[] = [];
  public isMoodleUserSignedIn: boolean;

  @ViewChild(LoginComponent) login: LoginComponent;

  // TODO remove when presentation is done
  private floVorlesungDeutsch = [
    {
      name: 'Prof. Helmut Holz',
      description: '<div class="no-overflow"><p>- Heute -</p><p>Vorlesung 27.02.2021 11:00</p><p>(Prüfung: 23.04.2021)</p><p>Vertiefungsrichtung: Softwareentwicklung</p></div>',
      location: 'KA.206',
      id: 1,
      descriptionformat: 1,
      categoryid: null,
      groupid: null,
      userid: 1,
      repeatid: null,
      eventcount: null,
      modulename: '',
      instance: null,
      eventtype: '',
      timestart: 1,
      timeduration: 1,
      timesort: 1,
      visible: 1,
      timemodified: 1,
      icon: {
        key: '',
        component: '',
        alttext: ''
      },
      course: {
        id: 1,
        fullname: 'Objektorientierte Software-Entwicklung (Vertiefungsmodul 1)',
        shortname: '',
        idnumber: '',
        summary: '',
        summaryformat: 1,
        startdate: 1,
        enddate: 1,
        visible: false,
        fullnamedisplay: '',
        viewurl: '',
        courseimage: '',
        progress: 1,
        hasprogress: false,
        isfavourite: false,
        hidden: false,
        showshortname: false,
        coursecategory: '',
      },
      category: {
        id: 1,
        name: '',
        idnumber: '',
        parent: 1,
        coursecount: 1,
        visible: 1,
        timemodified: 1,
        depth: 1,
        nestedname: '',
        url: '',
      },
      subscription: {
        displayeventsource: false,
        subscriptionname: '',
        subscriptionurl: '',
      },
      canedit: false,
      candelete: false,
      deleteurl: '',
      editurl: '',
      viewurl: '',
      formattedtime: '',
      isactionevent: false,
      iscourseevent: false,
      iscategoryevent: false,
      groupname: null,
      normalisedeventtype: '',
      normalisedeventtypetext: '',
      url: '',
      islastday: false,
      popupname: '',
      draggable: false
    },
    {
      name: 'Dipl.-Ing. Matthias Roth',
      description: '<div class="no-overflow"><p>- Heute -</p><p>Vorlesung 27.02.2021 14.00</p><p>(Prüfung: 18.06.2021)</p><p>Vertiefungsrichtung: Softwareentwicklung</p></div>',
      location: 'KA.309',
      id: 1,
      descriptionformat: 1,
      categoryid: null,
      groupid: null,
      userid: 1,
      repeatid: null,
      eventcount: null,
      modulename: '',
      instance: null,
      eventtype: '',
      timestart: 1,
      timeduration: 1,
      timesort: 1,
      visible: 1,
      timemodified: 1,
      icon: {
        key: '',
        component: '',
        alttext: ''
      },
      course: {
        id: 1,
        fullname: 'Datenbankentwicklung und Webtechnologien (Vertiefungsmodul 2)',
        shortname: '',
        idnumber: '',
        summary: '',
        summaryformat: 1,
        startdate: 1,
        enddate: 1,
        visible: false,
        fullnamedisplay: '',
        viewurl: '',
        courseimage: '',
        progress: 1,
        hasprogress: false,
        isfavourite: false,
        hidden: false,
        showshortname: false,
        coursecategory: '',
      },
      category: {
        id: 1,
        name: '',
        idnumber: '',
        parent: 1,
        coursecount: 1,
        visible: 1,
        timemodified: 1,
        depth: 1,
        nestedname: '',
        url: '',
      },
      subscription: {
        displayeventsource: false,
        subscriptionname: '',
        subscriptionurl: '',
      },
      canedit: false,
      candelete: false,
      deleteurl: '',
      editurl: '',
      viewurl: '',
      formattedtime: '',
      isactionevent: false,
      iscourseevent: false,
      iscategoryevent: false,
      groupname: null,
      normalisedeventtype: '',
      normalisedeventtypetext: '',
      url: '',
      islastday: false,
      popupname: '',
      draggable: false
    },
    {
      name: 'Prof. Dr. Matthias Hopf',
      description: '<div class="no-overflow"><p>- Montag -</p><p>Vorlesung 29.02.2021 10.00</p><p>(Prüfung: 19.09.2021)</p><p>Vertiefungsrichtung: Softwareentwicklung</p></div>',
      location: 'KA.234',
      id: 1,
      descriptionformat: 1,
      categoryid: null,
      groupid: null,
      userid: 1,
      repeatid: null,
      eventcount: null,
      modulename: '',
      instance: null,
      eventtype: '',
      timestart: 1,
      timeduration: 1,
      timesort: 1,
      visible: 1,
      timemodified: 1,
      icon: {
        key: '',
        component: '',
        alttext: ''
      },
      course: {
        id: 1,
        fullname: 'Computergrafik (Vertiefungsmodul 4)',
        shortname: '',
        idnumber: '',
        summary: '',
        summaryformat: 1,
        startdate: 1,
        enddate: 1,
        visible: false,
        fullnamedisplay: '',
        viewurl: '',
        courseimage: '',
        progress: 1,
        hasprogress: false,
        isfavourite: false,
        hidden: false,
        showshortname: false,
        coursecategory: '',
      },
      category: {
        id: 1,
        name: '',
        idnumber: '',
        parent: 1,
        coursecount: 1,
        visible: 1,
        timemodified: 1,
        depth: 1,
        nestedname: '',
        url: '',
      },
      subscription: {
        displayeventsource: false,
        subscriptionname: '',
        subscriptionurl: '',
      },
      canedit: false,
      candelete: false,
      deleteurl: '',
      editurl: '',
      viewurl: '',
      formattedtime: '',
      isactionevent: false,
      iscourseevent: false,
      iscategoryevent: false,
      groupname: null,
      normalisedeventtype: '',
      normalisedeventtypetext: '',
      url: '',
      islastday: false,
      popupname: '',
      draggable: false
    }
  ];

  // TODO remove when presentation is done
  private semjonVorlesungEnglisch = [
    {
      name: 'Christian S. Fötinger',
      description: '<div class="no-overflow"><p>- Heute -</p><p>Vorlesung 27.02.2021 10:30</p><p>(Prüfung: 18.06.2021)</p><p>Vertiefungsrichtung: IT-Security Engineering</p></div>',
      location: 'KA.200',
      id: 1,
      descriptionformat: 1,
      categoryid: null,
      groupid: null,
      userid: 1,
      repeatid: null,
      eventcount: null,
      modulename: '',
      instance: null,
      eventtype: '',
      timestart: 1,
      timeduration: 1,
      timesort: 1,
      visible: 1,
      timemodified: 1,
      icon: {
        key: '',
        component: '',
        alttext: ''
      },
      course: {
        id: 1,
        fullname: 'Governance, Frameworks & Standards (Vertiefungsmodul 2)',
        shortname: '',
        idnumber: '',
        summary: '',
        summaryformat: 1,
        startdate: 1,
        enddate: 1,
        visible: false,
        fullnamedisplay: '',
        viewurl: '',
        courseimage: '',
        progress: 1,
        hasprogress: false,
        isfavourite: false,
        hidden: false,
        showshortname: false,
        coursecategory: '',
      },
      category: {
        id: 1,
        name: '',
        idnumber: '',
        parent: 1,
        coursecount: 1,
        visible: 1,
        timemodified: 1,
        depth: 1,
        nestedname: '',
        url: '',
      },
      subscription: {
        displayeventsource: false,
        subscriptionname: '',
        subscriptionurl: '',
      },
      canedit: false,
      candelete: false,
      deleteurl: '',
      editurl: '',
      viewurl: '',
      formattedtime: '',
      isactionevent: false,
      iscourseevent: false,
      iscategoryevent: false,
      groupname: null,
      normalisedeventtype: '',
      normalisedeventtypetext: '',
      url: '',
      islastday: false,
      popupname: '',
      draggable: false
    },
    {
      name: 'Gregor Heilmeier',
      description: '<div class="no-overflow"><p>- Heute -</p><p>Vorlesung 27.02.2021 13.15</p><p>(Prüfung: 23.04.2021)</p><p>Vertiefungsrichtung: Digitalisierung</p></div>',
      location: 'KA.113',
      id: 1,
      descriptionformat: 1,
      categoryid: null,
      groupid: null,
      userid: 1,
      repeatid: null,
      eventcount: null,
      modulename: '',
      instance: null,
      eventtype: '',
      timestart: 1,
      timeduration: 1,
      timesort: 1,
      visible: 1,
      timemodified: 1,
      icon: {
        key: '',
        component: '',
        alttext: ''
      },
      course: {
        id: 1,
        fullname: 'Organisationsentwicklung (Vertiefungsmodul 4)',
        shortname: '',
        idnumber: '',
        summary: '',
        summaryformat: 1,
        startdate: 1,
        enddate: 1,
        visible: false,
        fullnamedisplay: '',
        viewurl: '',
        courseimage: '',
        progress: 1,
        hasprogress: false,
        isfavourite: false,
        hidden: false,
        showshortname: false,
        coursecategory: '',
      },
      category: {
        id: 1,
        name: '',
        idnumber: '',
        parent: 1,
        coursecount: 1,
        visible: 1,
        timemodified: 1,
        depth: 1,
        nestedname: '',
        url: '',
      },
      subscription: {
        displayeventsource: false,
        subscriptionname: '',
        subscriptionurl: '',
      },
      canedit: false,
      candelete: false,
      deleteurl: '',
      editurl: '',
      viewurl: '',
      formattedtime: '',
      isactionevent: false,
      iscourseevent: false,
      iscategoryevent: false,
      groupname: null,
      normalisedeventtype: '',
      normalisedeventtypetext: '',
      url: '',
      islastday: false,
      popupname: '',
      draggable: false
    },
    {
      name: 'Andreas Rothlauf',
      description: '<div class="no-overflow"><p>- Montag -</p><p>Vorlesung 29.02.2021 09.30</p><p>(Prüfung: 28.11.2021)</p></div>',
      location: 'KA.311',
      id: 1,
      descriptionformat: 1,
      categoryid: null,
      groupid: null,
      userid: 1,
      repeatid: null,
      eventcount: null,
      modulename: '',
      instance: null,
      eventtype: '',
      timestart: 1,
      timeduration: 1,
      timesort: 1,
      visible: 1,
      timemodified: 1,
      icon: {
        key: '',
        component: '',
        alttext: ''
      },
      course: {
        id: 1,
        fullname: 'Einführung in Continuous Integration, Continuous Delivery und Deployment',
        shortname: '',
        idnumber: '',
        summary: '',
        summaryformat: 1,
        startdate: 1,
        enddate: 1,
        visible: false,
        fullnamedisplay: '',
        viewurl: '',
        courseimage: '',
        progress: 1,
        hasprogress: false,
        isfavourite: false,
        hidden: false,
        showshortname: false,
        coursecategory: '',
      },
      category: {
        id: 1,
        name: '',
        idnumber: '',
        parent: 1,
        coursecount: 1,
        visible: 1,
        timemodified: 1,
        depth: 1,
        nestedname: '',
        url: '',
      },
      subscription: {
        displayeventsource: false,
        subscriptionname: '',
        subscriptionurl: '',
      },
      canedit: false,
      candelete: false,
      deleteurl: '',
      editurl: '',
      viewurl: '',
      formattedtime: '',
      isactionevent: false,
      iscourseevent: false,
      iscategoryevent: false,
      groupname: null,
      normalisedeventtype: '',
      normalisedeventtypetext: '',
      url: '',
      islastday: false,
      popupname: '',
      draggable: false
    }
  ];

  // TODO remove when presentation is done
  private moodleUserName: string;

  async ionViewWillEnter() {
    await this.login.checkMoodleLoginState();
  }

  public async onSignIn(isSignedIn: boolean) {
    isSignedIn ? this.isMoodleUserSignedIn = true : this.isMoodleUserSignedIn = false;
  }

  public async fetchMoodleData(moodleToken: MoodleToken) {
    const loading = await this.loadingController.create({
      message: 'Collecting moodle data...'
    });
    await loading.present();

    this.token = moodleToken;
    const calenderRequestData = await this.moodleService.getCalenderEventsWeek(moodleToken).toPromise();

    if (this.moodleService.containsEvents(calenderRequestData)) {
      // remove images => needs cookies
      // TODO remove when presentation is done
      await this.storage.get('moodle_user').then(userName => {
        this.moodleUserName = userName;
      });
      // this.calenderEvents = this.CleanupEvents(calenderRequestData.events);
      this.calenderEvents = this.moodleUserName === 'BarthFl81369' ? this.floVorlesungDeutsch : this.semjonVorlesungEnglisch;

      // add dummy location to KA.206
      // TODO uncomment when presentation is done
      // for (const event of this.calenderEvents) {
      //   event.location = 'KA.206';
      // }
      await loading.dismiss();
    } else {
      this.isMoodleUserSignedIn = false;
      this.login.isUserLoggedIn = false;
      await loading.dismiss();
    }
  }

  public async doRefreshEvents(event) {
    if (this.isMoodleUserSignedIn) {
      const calenderRequestData = await this.moodleService.getCalenderEventsWeek(this.token).toPromise();
      if (this.moodleService.containsEvents(calenderRequestData)) {
        // TODO remove when presentation is done
        await this.storage.get('moodle_user').then(userName => {
          this.moodleUserName = userName;
        });
        // this.calenderEvents = calenderRequestData.events;
        this.calenderEvents = this.moodleUserName === 'BarthFl81369' ? this.floVorlesungDeutsch : this.semjonVorlesungEnglisch;

        // add dummy location to KA.206
        // TODO uncomment when presentation is done
        // for (const calenderEvent of this.calenderEvents) {
        //   calenderEvent.location = 'KA.206';
        // }

        event.target.complete();
      } else {
        this.isMoodleUserSignedIn = false;
      }
    }
  }

  public async onLocationClick(locationName: string) {
    await this.router.navigate(['tabs/navigation/detail'], { queryParams: { location: locationName } });
  }

  ngAfterViewInit(): void {}

   private CleanupEvents(events: Event[]) {
     const imgRegex = new RegExp('<img[^>]*?>', 'g');
     for (const event of events) {
        if (imgRegex.test(event.description)) {
          for (const match of event.description.match(imgRegex)) {
            event.description = event.description.replace(match, '');
          }
        }
     }

     return events;
  }
}
