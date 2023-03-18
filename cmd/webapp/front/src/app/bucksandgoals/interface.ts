interface bucks{
    anz: number
    cba: number
    goal: number
}

export interface data{
    time: string
    data: bucks
    savings: number
}

export interface ChartDataset{
    label: string
    backgroundColor: string
    borderColor: string
    data: number[]
    borderDash: number[]
}

export interface ChartData{
    labels: string []
    datasets: ChartDataset []
}