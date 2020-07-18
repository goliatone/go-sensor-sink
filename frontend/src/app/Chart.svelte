<script>
    import {onMount, afterUpdate} from 'svelte';
    import moment from 'moment';
    import Chart from 'chart.js';

    export let data;

    const chartColors = {
        red: 'rgb(255, 99, 132)',
        orange: 'rgb(255, 159, 64)',
        yellow: 'rgb(255, 205, 86)',
        green: 'rgb(75, 192, 192)',
        blue: 'rgb(54, 162, 235)',
        purple: 'rgb(153, 102, 255)',
        grey: 'rgb(201, 203, 207)'
    };

    const colorNames = Object.keys(chartColors);

    let ctx;
    let myPie;

    function randomScalingFactor() {
        return Math.round(Math.random() * 100);
    }
    
    function newDate(days) {
        return moment().add(days, 'd').toDate();
    }

    function newDateString(days) {
        return moment().add(days, 'd').format();
    }

    let config = {
        type: 'line',
        data: {
            datasets: [{
                label: 'Dataset with string point data',
                borderColor: chartColors.red,
                backgroundColor: chartColors.orange,
                fill: false,
                data: [{
                    x: newDateString(0),
                    y: randomScalingFactor()
                }, {
                    x: newDateString(2),
                    y: randomScalingFactor()
                }, {
                    x: newDateString(4),
                    y: randomScalingFactor()
                }, {
                    x: newDateString(5),
                    y: randomScalingFactor()
                }],
            }, {
                label: 'Dataset with date object point data',
                borderColor: chartColors.blue,
                backgroundColor: chartColors.purple,
                fill: false,
                data: [{
                    x: newDate(0),
                    y: randomScalingFactor()
                }, {
                    x: newDate(2),
                    y: randomScalingFactor()
                }, {
                    x: newDate(4),
                    y: randomScalingFactor()
                }, {
                    x: newDate(5),
                    y: randomScalingFactor()
                }]
            }]
        },
        options: {
            responsive: true,
            title: {
                display: true,
                text: 'Humidity & Temperature'
            },
            
            scales: {
                
                xAxes: [{
                    type: 'time',
                    display: true,
                    scaleLabel: {
                        display: true,
                        labelString: 'Date'
                    },
                    ticks: {
                        major: {
                            fontStyle: 'bold',
                            fontColor: '#FF0000'
                        }
                    },
                    gridLines: {
                        display: true,
                        color: 'rgba(224,224,224,0.1)',
                        z: -1
                    },
                }],
                yAxes: [{
                    display: true,
                    scaleLabel: {
                        display: true,
                        labelString: 'value'
                    },
                    gridLines: {
                        display: true,
                        color: 'rgba(224,224,224,0.1)',
                        z: -1
                    },
                }]
            }
        }
    };

    window.config = config;

    onMount(_=>{
        ctx = document.getElementById('chart-area').getContext('2d');
    });

    function initializeChart() {
        console.log('after update?')
    
        /**
         * Attach the graph to the DOM
         */ 
        
        
        myPie = new Chart(ctx, config);
    }

    /**
     * Handle button click for randomized dataset
     */ 
    function randomizeData() {
        config.data.datasets.forEach(dataset => {
            dataset.data.forEach(dataObj => {
                dataObj.y = randomScalingFactor();
            });
        });

        myPie.update();
    }
    /**
     * Handle adding a new dataset
     */ 
    function addDataset() {
        if (config.data.datasets.length > 0) {
            config.data.datasets[0].data.push({
                x: newDateString(config.data.datasets[0].data.length + 2),
                y: randomScalingFactor()
            });
            config.data.datasets[1].data.push({
                x: newDate(config.data.datasets[1].data.length + 2),
                y: randomScalingFactor()
            });

            myPie.update();
        }
    }

    /**
     * Handle removing dataset
     */ 
    function removeDataset(){
        // config.data.datasets.splice(0, 1);
        config.data.datasets.forEach(function(dataset) {
            dataset.data.pop();
        });
        myPie.update();
    }

    afterUpdate(initializeChart);
</script>

<div id="canvas-holder" style="width:3;height=1;" >
    <canvas id="chart-area"></canvas>
</div>

<button id="randomizeData" on:click={_=> randomizeData()}>Randomize Data</button>
<button id="addDataset" on:click={_=> addDataset()}>Add Dataset</button>
<button id="removeDataset" on:click={_=> removeDataset()}>Remove Dataset</button>