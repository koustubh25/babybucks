import { Injectable } from '@angular/core';
import { ChartConfiguration, ChartDataset } from 'chart.js';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';
import { data } from './interface';

@Injectable({
  providedIn: 'root'
})
export class BucksandgoalsService {

  constructor(private http: HttpClient) { }

  getData(): Observable<data[]>{
    return this.http.get<data[]>("http://192.168.18.8:3000/fetchbalancesandgoals", {
      headers: {'Access-Control-Allow-Origin': '*' }
    })
  }

  // getChartConfigurationData(): Observable<ChartConfiguration["data"]>  {
    


  //   const data = {
  //     labels: ['January','February','March','April','May'],
  //     datasets: [{
  //       label: 'My First dataset',
  //       backgroundColor: 'rgb(255, 99, 132)',
  //       borderColor: 'rgb(255, 99, 132)',
  //       data: [10, 5, 2, 20, 30, 45],
  //     }]
  // }
    // return data

  // }
}
