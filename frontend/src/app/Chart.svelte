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

    let ctx;
    let chart;
    let unit;

    // 631180800009-
    // 1594793774---
    // 1594793774002

    var readings = [
        {
            "time": "2020-07-15T06:20:31.457047Z",
            "id": "esp01",
            "h": 54.8,
            "t": 25.7
        },
        {
            "time": "2020-07-15T06:20:28.775896Z",
            "id": "esp01",
            "h": 54.9,
            "t": 25.7
        },
        {
            "time": "2020-07-15T06:20:26.384234Z",
            "id": "esp01",
            "h": 54.9,
            "t": 25.7
        },
        {
            "time": "2020-07-15T06:20:24.030188Z",
            "id": "esp01",
            "h": 54.9,
            "t": 25.7
        },
        {
            "time": "2020-07-15T06:20:21.77513Z",
            "id": "esp01",
            "h": 54.9,
            "t": 25.7
        },
        {
            "time": "2020-07-15T06:20:18.367697Z",
            "id": "esp01",
            "h": 54.8,
            "t": 25.7
        },
        {
            "time": "2020-07-15T06:20:16.349251Z",
            "id": "esp01",
            "h": 54.8,
            "t": 25.7
        },
        {
            "time": "2020-07-15T06:20:13.696202Z",
            "id": "esp01",
            "h": 54.8,
            "t": 25.7
        },
        {
            "time": "2020-07-15T06:20:10.505013Z",
            "id": "esp01",
            "h": 54.8,
            "t": 25.7
        },
        {
            "time": "2020-07-15T06:20:07.895129Z",
            "id": "esp01",
            "h": 54.8,
            "t": 25.7
        },
        {
            "time": "2020-07-15T06:20:06.313558Z",
            "id": "esp01",
            "h": 54.8,
            "t": 25.7
        },
        {
            "time": "2020-07-15T06:20:04.516541Z",
            "id": "esp01",
            "h": 54.8,
            "t": 25.7
        },
        {
            "time": "2020-07-15T06:20:00.989231Z",
            "id": "esp01",
            "h": 54.8,
            "t": 25.7
        },
        {
            "time": "2020-07-15T06:19:57.846061Z",
            "id": "esp01",
            "h": 54.7,
            "t": 25.7
        },
        {
            "time": "2020-07-15T06:19:55.151575Z",
            "id": "esp01",
            "h": 54.7,
            "t": 25.7
        },
        {
            "time": "2020-07-15T06:19:52.386824Z",
            "id": "esp01",
            "h": 54.7,
            "t": 25.7
        },
        {
            "time": "2020-07-15T06:19:49.929024Z",
            "id": "esp01",
            "h": 54.7,
            "t": 25.7
        },
        {
            "time": "2020-07-15T06:19:47.164675Z",
            "id": "esp01",
            "h": 54.7,
            "t": 25.7
        },
        {
            "time": "2020-07-15T06:19:44.520155Z",
            "id": "esp01",
            "h": 54.7,
            "t": 25.7
        },
        {
            "time": "2020-07-15T06:19:41.941839Z",
            "id": "esp01",
            "h": 54.6,
            "t": 25.7
        },
        {
            "time": "2020-07-15T06:19:39.382049Z",
            "id": "esp01",
            "h": 54.6,
            "t": 25.7
        },
        {
            "time": "2020-07-15T06:19:36.577095Z",
            "id": "esp01",
            "h": 54.6,
            "t": 25.7
        },
        {
            "time": "2020-07-15T06:19:34.519844Z",
            "id": "esp01",
            "h": 54.6,
            "t": 25.7
        },
        {
            "time": "2020-07-15T06:19:31.394782Z",
            "id": "esp01",
            "h": 54.6,
            "t": 25.7
        },
        {
            "time": "2020-07-15T06:19:28.766426Z",
            "id": "esp01",
            "h": 54.6,
            "t": 25.7
        },
        {
            "time": "2020-07-15T06:19:26.172221Z",
            "id": "esp01",
            "h": 54.6,
            "t": 25.7
        },
        {
            "time": "2020-07-15T06:19:23.55384Z",
            "id": "esp01",
            "h": 54.7,
            "t": 25.7
        },
        {
            "time": "2020-07-15T06:19:21.046692Z",
            "id": "esp01",
            "h": 54.7,
            "t": 25.7
        },
        {
            "time": "2020-07-15T06:19:18.300664Z",
            "id": "esp01",
            "h": 54.7,
            "t": 25.7
        },
        {
            "time": "2020-07-15T06:19:16.039019Z",
            "id": "esp01",
            "h": 54.6,
            "t": 25.7
        },
        {
            "time": "2020-07-15T06:19:13.269815Z",
            "id": "esp01",
            "h": 54.5,
            "t": 25.7
        },
        {
            "time": "2020-07-15T06:19:10.473045Z",
            "id": "esp01",
            "h": 54.4,
            "t": 25.7
        },
        {
            "time": "2020-07-15T06:19:07.866211Z",
            "id": "esp01",
            "h": 54.5,
            "t": 25.7
        },
        {
            "time": "2020-07-15T06:19:05.282523Z",
            "id": "esp01",
            "h": 54.5,
            "t": 25.7
        },
        {
            "time": "2020-07-15T06:19:02.927052Z",
            "id": "esp01",
            "h": 54.5,
            "t": 25.7
        },
        {
            "time": "2020-07-15T06:19:00.264703Z",
            "id": "esp01",
            "h": 54.5,
            "t": 25.7
        },
        {
            "time": "2020-07-15T06:18:57.845871Z",
            "id": "esp01",
            "h": 54.5,
            "t": 25.7
        },
        {
            "time": "2020-07-15T06:18:55.144953Z",
            "id": "esp01",
            "h": 54.5,
            "t": 25.7
        },
        {
            "time": "2020-07-15T06:18:52.380061Z",
            "id": "esp01",
            "h": 54.5,
            "t": 25.7
        },
        {
            "time": "2020-07-15T06:18:49.922387Z",
            "id": "esp01",
            "h": 54.4,
            "t": 25.7
        },
        {
            "time": "2020-07-15T06:18:47.017786Z",
            "id": "esp01",
            "h": 54.4,
            "t": 25.7
        },
        {
            "time": "2020-07-15T06:18:44.392622Z",
            "id": "esp01",
            "h": 54.4,
            "t": 25.7
        },
        {
            "time": "2020-07-15T06:18:41.935389Z",
            "id": "esp01",
            "h": 54.4,
            "t": 25.7
        },
        {
            "time": "2020-07-15T06:18:39.784684Z",
            "id": "esp01",
            "h": 53.6,
            "t": 25.7
        },
        {
            "time": "2020-07-15T06:18:36.729165Z",
            "id": "esp01",
            "h": 54.4,
            "t": 25.7
        },
        {
            "time": "2020-07-15T06:18:34.519212Z",
            "id": "esp01",
            "h": 54.4,
            "t": 25.7
        },
        {
            "time": "2020-07-15T06:18:31.798038Z",
            "id": "esp01",
            "h": 54.5,
            "t": 25.7
        },
        {
            "time": "2020-07-15T06:18:28.824511Z",
            "id": "esp01",
            "h": 54.4,
            "t": 25.7
        },
        {
            "time": "2020-07-15T06:18:26.267702Z",
            "id": "esp01",
            "h": 54.3,
            "t": 25.7
        },
        {
            "time": "2020-07-15T06:18:23.810003Z",
            "id": "esp01",
            "h": 54.3,
            "t": 25.7
        },
        {
            "time": "2020-07-15T06:18:21.045192Z",
            "id": "esp01",
            "h": 54.3,
            "t": 25.7
        },
        {
            "time": "2020-07-15T06:18:18.904636Z",
            "id": "esp01",
            "h": 54.3,
            "t": 25.7
        },
        {
            "time": "2020-07-15T06:18:15.822823Z",
            "id": "esp01",
            "h": 54.2,
            "t": 25.7
        },
        {
            "time": "2020-07-15T06:18:13.045265Z",
            "id": "esp01",
            "h": 54.2,
            "t": 25.7
        },
        {
            "time": "2020-07-15T06:18:10.433377Z",
            "id": "esp01",
            "h": 54.1,
            "t": 25.7
        },
        {
            "time": "2020-07-15T06:18:07.846046Z",
            "id": "esp01",
            "h": 54.1,
            "t": 25.7
        },
        {
            "time": "2020-07-15T06:18:05.378237Z",
            "id": "esp01",
            "h": 54.1,
            "t": 25.7
        },
        {
            "time": "2020-07-15T06:18:02.9202Z",
            "id": "esp01",
            "h": 54.1,
            "t": 25.7
        },
        {
            "time": "2020-07-15T06:18:00.155403Z",
            "id": "esp01",
            "h": 54.1,
            "t": 25.7
        },
        {
            "time": "2020-07-15T06:17:57.844796Z",
            "id": "esp01",
            "h": 54.2,
            "t": 25.7
        },
        {
            "time": "2020-07-15T06:17:54.932917Z",
            "id": "esp01",
            "h": 54.2,
            "t": 25.7
        },
        {
            "time": "2020-07-15T06:17:52.564797Z",
            "id": "esp01",
            "h": 54.2,
            "t": 25.7
        },
        {
            "time": "2020-07-15T06:17:50.018696Z",
            "id": "esp01",
            "h": 54.1,
            "t": 25.7
        },
        {
            "time": "2020-07-15T06:17:47.016692Z",
            "id": "esp01",
            "h": 54.1,
            "t": 25.7
        },
        {
            "time": "2020-07-15T06:17:44.517176Z",
            "id": "esp01",
            "h": 54.1,
            "t": 25.7
        },
        {
            "time": "2020-07-15T06:17:41.718202Z",
            "id": "esp01",
            "h": 54,
            "t": 25.7
        },
        {
            "time": "2020-07-15T06:17:39.10843Z",
            "id": "esp01",
            "h": 54,
            "t": 25.7
        },
        {
            "time": "2020-07-15T06:17:36.728012Z",
            "id": "esp01",
            "h": 54,
            "t": 25.7
        },
        {
            "time": "2020-07-15T06:17:34.350405Z",
            "id": "esp01",
            "h": 54.1,
            "t": 25.8
        },
        {
            "time": "2020-07-15T06:17:32.815162Z",
            "id": "esp01",
            "h": 54,
            "t": 25.7
        },
        {
            "time": "2020-07-15T06:17:28.73396Z",
            "id": "esp01",
            "h": 54,
            "t": 25.7
        },
        {
            "time": "2020-07-15T06:17:26.058768Z",
            "id": "esp01",
            "h": 54,
            "t": 25.7
        },
        {
            "time": "2020-07-15T06:17:23.447525Z",
            "id": "esp01",
            "h": 54,
            "t": 25.7
        },
        {
            "time": "2020-07-15T06:17:21.140897Z",
            "id": "esp01",
            "h": 54.1,
            "t": 25.8
        },
        {
            "time": "2020-07-15T06:17:18.990424Z",
            "id": "esp01",
            "h": 54.1,
            "t": 25.8
        },
        {
            "time": "2020-07-15T06:17:15.887876Z",
            "id": "esp01",
            "h": 54.1,
            "t": 25.7
        },
        {
            "time": "2020-07-15T06:17:13.032534Z",
            "id": "esp01",
            "h": 54.1,
            "t": 25.8
        },
        {
            "time": "2020-07-15T06:17:10.397233Z",
            "id": "esp01",
            "h": 54.2,
            "t": 25.8
        },
        {
            "time": "2020-07-15T06:17:07.844183Z",
            "id": "esp01",
            "h": 54.3,
            "t": 25.8
        },
        {
            "time": "2020-07-15T06:17:05.78096Z",
            "id": "esp01",
            "h": 54.3,
            "t": 25.8
        },
        {
            "time": "2020-07-15T06:17:02.588454Z",
            "id": "esp01",
            "h": 54.3,
            "t": 25.8
        },
        {
            "time": "2020-07-15T06:16:59.955162Z",
            "id": "esp01",
            "h": 54.4,
            "t": 25.8
        },
        {
            "time": "2020-07-15T06:16:57.845214Z",
            "id": "esp01",
            "h": 53.6,
            "t": 25.8
        },
        {
            "time": "2020-07-15T06:16:54.732737Z",
            "id": "esp01",
            "h": 54.4,
            "t": 25.8
        },
        {
            "time": "2020-07-15T06:16:52.366022Z",
            "id": "esp01",
            "h": 54.3,
            "t": 25.7
        },
        {
            "time": "2020-07-15T06:16:49.80634Z",
            "id": "esp01",
            "h": 54.4,
            "t": 25.8
        },
        {
            "time": "2020-07-15T06:16:47.041192Z",
            "id": "esp01",
            "h": 54.4,
            "t": 25.8
        },
        {
            "time": "2020-07-15T06:16:44.518302Z",
            "id": "esp01",
            "h": 54.3,
            "t": 25.7
        },
        {
            "time": "2020-07-15T06:16:42.126054Z",
            "id": "esp01",
            "h": 54.4,
            "t": 25.8
        },
        {
            "time": "2020-07-15T06:16:39.36217Z",
            "id": "esp01",
            "h": 54.3,
            "t": 25.7
        },
        {
            "time": "2020-07-15T06:16:36.486592Z",
            "id": "esp01",
            "h": 54.4,
            "t": 25.8
        },
        {
            "time": "2020-07-15T06:16:33.854706Z",
            "id": "esp01",
            "h": 54.3,
            "t": 25.7
        },
        {
            "time": "2020-07-15T06:16:31.715871Z",
            "id": "esp01",
            "h": 54.3,
            "t": 25.7
        },
        {
            "time": "2020-07-15T06:16:28.712432Z",
            "id": "esp01",
            "h": 54.3,
            "t": 25.7
        },
        {
            "time": "2020-07-15T06:16:26.091089Z",
            "id": "esp01",
            "h": 54.4,
            "t": 25.8
        },
        {
            "time": "2020-07-15T06:16:24.308524Z",
            "id": "esp01",
            "h": 54.4,
            "t": 25.8
        },
        {
            "time": "2020-07-15T06:16:20.930025Z",
            "id": "esp01",
            "h": 54.4,
            "t": 25.8
        },
        {
            "time": "2020-07-15T06:16:18.778743Z",
            "id": "esp01",
            "h": 54.3,
            "t": 25.7
        },
        {
            "time": "2020-07-15T06:16:16.01412Z",
            "id": "esp01",
            "h": 54.2,
            "t": 25.7
        },
        {
            "time": "2020-07-15T06:16:14.002669Z",
            "id": "esp01",
            "h": 54.3,
            "t": 25.7
        }
    ];

    function getReadingsData(readings, type='temp') {
        let prop = type.includes('temp') ? 't' : 'h';
        let data = readings.map(r=>{
            return {
                t: parseInt(moment(r.time).format('x')),
                y: r[prop]
            };
        });

        return data;
    }
    
    function generateData() {

        function unitLessThanDay() {
            return unit === 'second' || unit === 'minute' || unit === 'hour';
        }

        function beforeNineThirty(date) {
            return date.hour() < 9 || (date.hour() === 9 && date.minute() < 30);
        }

        // Returns true if outside 9:30am-4pm on a weekday
        function outsideMarketHours(date) {
            if (date.isoWeekday() > 5) return true;
            
            if (unitLessThanDay() && (beforeNineThirty(date) || date.hour() > 16)) {
                return true;
            }

            return false;
        }

        function randomNumber(min, max) {
            return Math.random() * (max - min) + min;
        }

        function randomBar(date, lastClose) {
            var open = randomNumber(lastClose * 0.95, lastClose * 1.05).toFixed(2);
            var close = randomNumber(open * 0.95, open * 1.05).toFixed(2);
            return {
                t: date.valueOf(),
                y: close
            };
        }

        var date = moment('Jan 01 1990', 'MMM DD YYYY');
        var now = moment();
        var data = [];
        var lessThanDay = unitLessThanDay();
        for (; data.length < 600 && date.isBefore(now); date = date.clone().add(1, unit).startOf(unit)) {
            if (outsideMarketHours(date)) {
                if (!lessThanDay || !beforeNineThirty(date)) {
                    date = date.clone().add(date.isoWeekday() >= 5 ? 8 - date.isoWeekday() : 1, 'day');
                }
                if (lessThanDay) {
                    date = date.hour(9).minute(30).second(0);
                }
            }
            data.push(randomBar(date, data.length > 0 ? data[data.length - 1].y : 30));
        }

        return data;
    }

    let config = {
        data: {
            datasets: [
            {
                label: 'Humidity',
                borderColor: chartColors.red,
                backgroundColor: chartColors.orange,
                data: getReadingsData(readings, 'humidity'),
                type: 'line',
                pointRadius: 0,
                fill: false,
                lineTension: 0,
                borderWidth: 2
            },
            {
                label: 'Temperature',
                borderColor: chartColors.blue,
                backgroundColor: chartColors.purple,
                data: getReadingsData(readings, 'temperature'),
                type: 'line',
                pointRadius: 0,
                fill: false,
                lineTension: 0,
                borderWidth: 2
            }]
        },
        options: {
            animation: {
                duration: 0
            },
            scales: {
                xAxes: [{
                    type: 'time',
                    distribution: 'linear',//data are spread according to their time (distances can vary)
                    distribution: 'series',//data are spread at the same distance from each other
                    // bounds: 'ticks',
                    bounds: 'data',
                    time: {
                        unit: 'second',
                        displayFormats: {
                            quarter: 'h:mm:ss a'
                        }
                    },
                    offset: true,
                    ticks: {
                        major: {
                            enabled: true,
                            fontStyle: 'bold'
                        },
                        source: 'data',
                        // source: 'labels',
                        // source: 'auto',
                        autoSkip: true,
                        autoSkipPadding: 75,
                        maxRotation: 0,
                        sampleSize: 100
                    },
                    afterBuildTicks: function(scale, ticks) {
                        var majorUnit = scale._majorUnit;
                        var firstTick = ticks[0];
                        var i, ilen, val, tick, currMajor, lastMajor;

                        val = moment(ticks[0].value);
                        if ((majorUnit === 'minute' && val.second() === 0)
                                || (majorUnit === 'hour' && val.minute() === 0)
                                || (majorUnit === 'day' && val.hour() === 9)
                                || (majorUnit === 'month' && val.date() <= 3 && val.isoWeekday() === 1)
                                || (majorUnit === 'year' && val.month() === 0)) {
                            firstTick.major = true;
                        } else {
                            firstTick.major = false;
                        }
                        lastMajor = val.get(majorUnit);

                        for (i = 1, ilen = ticks.length; i < ilen; i++) {
                            tick = ticks[i];
                            val = moment(tick.value);
                            currMajor = val.get(majorUnit);
                            tick.major = currMajor !== lastMajor;
                            lastMajor = currMajor;
                        }
                        return ticks;
                    }
                }],
                yAxes: [{
                    gridLines: {
                        drawBorder: false
                    },
                    scaleLabel: {
                        display: true,
                        labelString: 'Temperature'
                    }
                }]
            },
            tooltips: {
                intersect: false,
                mode: 'index',
                callbacks: {
                    label: function(tooltipItem, myData) {
                        var label = myData.datasets[tooltipItem.datasetIndex].label || '';
                        if (label) {
                            label += ': ';
                        }
                        label += parseFloat(tooltipItem.value).toFixed(2);
                        return label;
                    }
                }
            }
        }
    };

    window.config = config;
    window.moment = moment;
    window.getReadings = _=> getReadingsData(readings);

    onMount(_=>{
        unit = document.getElementById('unit').value;
        ctx = document.getElementById('chart1').getContext('2d');
        ctx.canvas.width = 1000;
        ctx.canvas.height = 300;
    });

    function initializeChart() {
    
        /**
         * Attach the graph to the DOM
         */ 
        
        
        chart = new Chart(ctx, config);
    }

    /**
     * Handle button click for randomized dataset
     */ 
    function update() {
        function updateDateset(index, type) {
            var type = document.getElementById('type').value;
            var dataset = chart.config.data.datasets[index];
            dataset.type = type;
            dataset.data = getReadingsData(readings, type);
            // dataset.data = generateData();
        }
        updateDateset(0, 'humitidy');
        updateDateset(1, 'temperature');

        chart.update();
    }
    

    afterUpdate(initializeChart);
</script>

<div id="canvas-holder" style="width:3;height=1;" >
    <canvas id="chart1"></canvas>
</div>

Chart Type:
<select id="type">
    <option value="line">Line</option>
    <option value="bar">Bar</option>
</select>
<select id="unit">
    <option value="second">Second</option>
    <option value="minute">Minute</option>
    <option value="hour">Hour</option>
    <option value="day" selected>Day</option>
    <option value="month">Month</option>
    <option value="year">Year</option>
</select>
<button id="update" on:click={_=>update()} >update</button>