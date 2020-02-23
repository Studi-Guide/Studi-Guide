import { room } from '../building-objects-if';
import { section } from '../building-objects-if';

export let testDataRooms:room[] = [
  {
    name: 'KA.304',
    sections: [
      {
        start: { x: 0, y: 0, z: 3 },
        end: { x: 100, y: 0, z: 3 }
      },
      {
        start: { x: 100, y: 0, z: 3 },
        end: { x: 100, y: 100, z: 3 }
      },
      {
        start: { x: 100, y: 100, z: 3 },
        end: { x: 0, y: 100, z: 3 }
      },
      {
        start: { x: 0, y: 100, z: 3 },
        end: { x: 0, y: 0, z: 3 }
      }
    ],
    alias: '',
    doors: [],
    fill: '#676d7d'
  }
];

/*
,
{
  name: 'KA.2',
      fill: '#F80',
    width: 100,
    height: 125,
    x: 0,
    y: 50
},
{
  name: 'KA.3',
      fill: '5AF',
    width: 100,
    height: 125,
    x: 100,
    y: 0
},
{
  name: 'KA.4',
      fill: '#888',
    width: 100,
    height: 125,
    x: 100,
    y: 125
},
{
  name: 'KA.5',
      fill: '#F01A1A',
    width: 125,
    height: 75,
    x: 200,
    y: 0
},*/
