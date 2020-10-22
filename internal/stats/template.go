package stats

var (
	layoutTemplate = `
{{define "layout"}}
<html>
	<head>
		<title>Micro Stats</title>
		<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.6/css/bootstrap.min.css" integrity="sha384-1q8mTJOASx8j1Au+a5WDVnPi2lkFfwwEAa8hDDdjZlpLegxhjVME1fgjWPGmkzs7" crossorigin="anonymous">
                <link href="https://fonts.googleapis.com/css?family=Source+Code+Pro&display=swap" rel="stylesheet">
                <style>
                  html, body {
                    font-family: 'Source Code Pro', monospace;
                  }
                  html a { color: #333333; }
                  .navbar .navbar-brand { color: #333333; font-weight: bold; font-size: 2.0em; }
                  .navbar-brand img { display: inline; }
                  #navBar, .navbar-toggle { margin-top: 15px; }
                  .icon-bar { background-color: #333333; }
                </style>
                <style>
		<style>
		{{ template "style" . }}
		</style>
	</head>
	<body>
	  <nav class="navbar">
	    <div class="container">
	      <div class="navbar-header">
                <a class="navbar-brand logo" href="/"><img src="data:image/png;base64,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" height=50px width=auto /> Stats</a>
	      </div>
	    </div>
	  </nav>
          <div class="container">
            <div class="row">
              <span class="pull-right update h6"></span>
	      <div class="col-sm-4">
                <h4>&nbsp;</h4>
                <table class="table table-bordered">
                  <caption>Info</caption>
                  <tbody>
                    <tr>
                      <th>Started</th>
                      <td class="started"></td>
                    </tr>
                    <tr>
                      <th>Uptime</th>
                      <td class="uptime"></td>
                    </tr>
                    <tr>
                      <th>Memory</th>
                      <td class="memory"></td>
                    </tr>
                    <tr>
                      <th>Threads</th>
                      <td class="threads"></td>
                    </tr>
                    <tr>
                      <th>GC</th>
                      <td class="gc"></td>
                    </tr>
                  </tbody>
                </table>

                <table class="table table-bordered">
                  <caption>Requests</caption>
                  <tbody>
                    <tr>
                      <th>Total</th>
                      <td class="total"></td>
                    </tr>
                    <tr>
                      <th>20x</th>
                      <td class="20x"></td>
                    </tr>
                    <tr>
                      <th>40x</th>
                      <td class="40x"></td>
                    </tr>
                    <tr>
                      <th>50x</th>
                      <td class="50x"></td>
                    </tr>
                  </tbody>
                </table>
	      </div>
	      <div class="col-sm-8">
                {{ template "content" . }}
              </div>
            </div>
          </div>
	  <script src="https://cdnjs.cloudflare.com/ajax/libs/jquery/2.1.4/jquery.min.js"></script>
	  <script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.6/js/bootstrap.min.js" integrity="sha384-0mSbJDEHialfmuBBQP6A4Qrprq5OVfW37PRR3j5ELqxss1yVqOtnepnHVP9aJ7xS" crossorigin="anonymous"></script>
	  <script type="text/javascript" src="https://cdnjs.cloudflare.com/ajax/libs/canvasjs/1.7.0/canvasjs.min.js"></script>
	  {{template "script" . }}
	</body>
</html>
{{end}}
{{ define "style" }}{{end}}
{{ define "script" }}{{end}}
{{ define "title" }}{{end}}
`

	statsTemplate = `
{{define "title"}}Stats{{end}}
{{define "content"}}
  <div id="chart" style="height: 300px; width: 100%;">
{{end}}
{{define "script"}}
<script>
  function loadChart(counters) {
	// dataPoints
	var dataPoints1 = [];
	var dataPoints2 = [];
	var dataPoints3 = [];

	var chart = new CanvasJS.Chart("chart",{
		zoomEnabled: true,
		title: {
			text: "Request Load"		
		},
		toolTip: {
			shared: true
			
		},
		legend: {
			verticalAlign: "top",
			horizontalAlign: "center",
			fontSize: 14,
			fontWeight: "bold",
			fontFamily: "calibri",
			fontColor: "dimGrey"
		},
		axisX: {
			title: "updates every 5 secs"
		},
		axisY:{
			includeZero: false
		}, 
		data: [{ 
			// dataSeries1
			type: "line",
			xValueType: "dateTime",
			showInLegend: true,
			name: "20x",
			dataPoints: dataPoints1
		},
		{				
			// dataSeries2
			type: "line",
			xValueType: "dateTime",
			showInLegend: true,
			name: "40x" ,
			dataPoints: dataPoints2
		},
		{				
			// dataSeries3
			type: "line",
			xValueType: "dateTime",
			showInLegend: true,
			name: "50x" ,
			dataPoints: dataPoints3
		}],
                legend:{
                cursor:"pointer",
                itemclick : function(e) {
                  if (typeof(e.dataSeries.visible) === "undefined" || e.dataSeries.visible) {
                    e.dataSeries.visible = false;
                  }
                  else {
                    e.dataSeries.visible = true;
                  }
                  chart.render();
                }
              }
	});

	var two = 0;
	var four = 0;
	var five = 0;

	for (i = 0; i < counters.length; i++) {
		var time = new Date((counters[i].timestamp + 5) * 1000);
                var counter = counters[i];

                if (counter["status_codes"]["20x"] != undefined) {
		  two = counter["status_codes"]["20x"];
                } else {
                  two = 0;
                }

                if (counter["status_codes"]["50x"] != undefined) {
		  five = counter["status_codes"]["50x"];
                } else {
                  five = 0;
                }

                if (counter["status_codes"]["40x"] != undefined) {
		  four = counter["status_codes"]["40x"];
                } else {
                  four = 0;
                }

		// pushing the new values
		dataPoints1.push({
			x: time.getTime(),
			y: two
		});
		dataPoints2.push({
			x: time.getTime(),
			y: four
		});
		dataPoints3.push({
			x: time.getTime(),
			y: five
		});
	}

	// updating legend text with  updated with y Value 
	chart.options.data[0].legendText = " 20x  " + two;
	chart.options.data[1].legendText = " 40x  " + four; 
	chart.options.data[2].legendText = " 50x  " + five;
	chart.render();
  };


  function loadStats() {
    var req = new XMLHttpRequest();
    req.onreadystatechange = function() {
	if (req.readyState == 4 && req.status == 200) {
	    console.log(req.responseText);

            var data = JSON.parse(req.responseText);
            var started = new Date(data["started"]*1000);
            var uptime = new Date() - started;

            // uptime
            uptime = uptime / 1000;
            if (uptime > 3600) {
              var time = uptime;
	      var hours   = Math.floor(time / 3600);
	      var minutes = Math.floor((time - (hours * 3600)) / 60);
	      var seconds = Math.floor(time - (hours * 3600) - (minutes * 60));

	      if (hours   < 10) {hours   = "0"+hours;}
	      if (minutes < 10) {minutes = "0"+minutes;}
	      if (seconds < 10) {seconds = "0"+seconds;}
	      uptime = hours+':'+minutes+':'+seconds;
            } else {
              uptime = uptime + "s";
            }

            // info
            $('.update').text("Last updated " + (new Date()).toUTCString());
            $('.started').text(started.toUTCString());
            $('.uptime').text(uptime);
            $('.memory').text(data["memory"]);
            $('.threads').text(data["threads"]);
            $('.gc').text(data["gc_pause"]);

            // requests
            var total = 0;
            var tx = 0;
            var fx = 0;
            var fox = 0;

            for (i = 0;  i < data["counters"].length; i++) {
              var counter = data["counters"][i];
              total += counter["total_reqs"];
              if (counter["status_codes"]["20x"] != undefined) {
                tx += counter["status_codes"]["20x"];
              };
              if (counter["status_codes"]["40x"] != undefined) {
                fox += counter["status_codes"]["40x"];
              };
              if (counter["status_codes"]["50x"] != undefined) {
                fx += counter["status_codes"]["50x"];
              };
            };

            $('.total').text(total);
            $('.20x').text(tx);
            $('.40x').text(fox);
            $('.50x').text(fx);

            loadChart(data["counters"]);
	}
    }

    var request = {};
    req.open("GET", window.location.href, true);
    req.setRequestHeader("Content-type","application/json");				
    req.send(JSON.stringify(request));

    setTimeout(function() {
      loadStats();
    }, 5000);
  };

  loadStats();
</script>
{{end}}
`
)
