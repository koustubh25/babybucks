import { Component,ElementRef, ViewChild } from '@angular/core';
import {Chart, ChartConfiguration, ChartItem, registerables, TimeScale} from 'chart.js'
import { timestamp } from 'rxjs';
import { BucksandgoalsService } from './bucksandgoals.service';
import { ChartData, ChartDataset } from './interface';

@Component({
  selector: 'app-bucksandgoals',
  templateUrl: './bucksandgoals.component.html',
  styleUrls: ['./bucksandgoals.component.css']
})


export class BucksandgoalsComponent {

  @ViewChild("mychart", {static: false}) span:ElementRef;

  ngAfterViewInit(): void {

    const chartDataSetSavings = {
      label: "",
      backgroundColor: "",
      borderColor: "",
      data: [],
      borderDash: []
    } as ChartDataset;
    const chartDataSetGoal = {
      label: "",
      backgroundColor: "",
      borderColor: "",
      data: [],
      borderDash: [10,5]
    } as ChartDataset

    const chartData = {
      labels: [],
      datasets: []
    } as ChartData;
    

    const testData = [{
      "time": "2023-03-13T02:57:40.376335Z",
      "data": {
        "anz": 21,
        "cba": 0,
        "goal": 23
      },
      "savings": 21
    },
    {
      "time": "2023-03-14T02:57:40.43812Z",
      "data": {
        "anz": 25,
        "cba": 0,
        "goal": 23
      },
      "savings": 25
    }];

    // this.bucksandgoalsService.getData().subscribe(d => {
      testData.forEach((el) => {
      // d.forEach(el => {
        chartData.labels.push(el.time);

        // savings
        chartDataSetSavings.backgroundColor = 'rgb(39, 120, 29)',
        chartDataSetSavings.borderColor = 'rgb(39, 120, 29)',
        chartDataSetSavings.data.push(el.savings)
        chartDataSetSavings.label = 'Savings'

        //goals
        chartDataSetGoal.backgroundColor = 'rgb(5, 223, 247)',
        chartDataSetGoal.borderColor = 'rgb(5, 223, 247)',
        chartDataSetGoal.data.push(el.data.goal)
        chartDataSetGoal.label = 'Goal'

      })
      console.log(chartData);
      chartData.datasets.push(chartDataSetSavings);
        chartData.datasets.push(chartDataSetGoal);
      
      this.createChart(chartData)
    // });
    
}

createChart(chartData: ChartData): void {
  Chart.register(TimeScale, ...registerables);
  
  const options = {
    scales: {
    //   x: {
    //     type: 'time',
    //     time: {
    //         displayFormats: {
    //             quarter: 'MMM YYYY'
    //         }
    //     }
    // }
    },
    plugins:{
    "legend": {
      "display": true,
      "position": "right" as const ,
      "align": "center" as const
    }
  },
    responsive: true,
  
}

  const config: ChartConfiguration = {
    type: 'line',
    data: chartData,
    options: options
  }

   const chartItem: ChartItem = this.span.nativeElement as ChartItem;

  new Chart(chartItem, config)
  }
  
  constructor(private el: ElementRef, private bucksandgoalsService: BucksandgoalsService) {
  }
 
}
