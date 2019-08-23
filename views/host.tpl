 {{define "host"}}
<div style="width:20%; float: right;">
    <div class="list-group">
        <a href="" class="list-group-item list-group-item-action active">Machine Info</a>
        <a class="list-group-item list-group-item-action" id="Hostname">Hostname: {{ .hostinfo.HostName }}</a>
        <a class="list-group-item list-group-item-action" id="HostOS">OS: {{ .hostinfo.OS }}</a>
        <a href="" class="list-group-item list-group-item-action active">CPU Info</a>
        <a class="list-group-item list-group-item-action" id="CPUModel">CPU Model: {{ .hostinfo.CPU.CPUModelandFrequency }}</a>
        <div id="accordion">
            <div class="card ">
                <div class="card-header" id="headingOne">
                    <h5 class="mb-0">
                        <button class="btn btn-link" data-toggle="collapse" data-target="#collapseOne" aria-expanded="true" aria-controls="collapseOne">CPU Cores Number: {{ .hostinfo.CPU.CPUCores }} & Core Utilizations</button>
                    </h5>
                </div>
                <div id="collapseOne" class="collapse" aria-labelledby="headingOne" data-parent="#accordion">
                    <div id="CPUpercentage">
                        <a href="#" class="list-group-item list-group-item-action"></a>
                    </div>
                </div>
            </div>
        </div>
        <a href="" class="list-group-item list-group-item-action active">Memory Info: </a>
        <a class="list-group-item list-group-item-action" id="TotalMemory">Total Memory: {{ .hostinfo.Memory.TotalMemory }} KB</a>
        <a class="list-group-item list-group-item-action" id="UsedMemory">Used Memory: {{ .hostinfo.Memory.UsedMemory }} KB</a>
        <a class="list-group-item list-group-item-action" id="MemoryUsage">Memory Usage: {{ .hostinfo.Memory.MemoryPercentage }} </a>
        <a class="list-group-item list-group-item-action" id="SWAP">Swap On or Off: {{ .hostinfo.Memory.SWAPonoff }}</a>
        <a href="" class="list-group-item list-group-item-action active">Disk Info</a>
        <a class="list-group-item list-group-item-action" id="TotalDisk">Total Disk: {{ .hostinfo.DiskSpace.TotalDisk }} B</a>
        <a class="list-group-item list-group-item-action" id="UsedDisk">Used Disk: {{ .hostinfo.DiskSpace.UsedDisk }} B</a>
        <a class="list-group-item list-group-item-action" id="AvailbleDisk">Availble Disk: {{ .hostinfo.DiskSpace.AvaileDisk }} B</a>
        <a class="list-group-item list-group-item-action" id="UsedDiskPercentage">Used Disk Percentage: {{ .hostinfo.DiskSpace.DiskPercentage }} %</a>
    </div>
    <form method="post" action="/host" style="text-align: center;">
        <input style="width:40%" type="button" value="Refresh" class="btn btn-primary mb-2" onclick="javascript:window.location.href='/host/';"> {{if .switch}}
        <input style="width:40%" type="submit" name="syncoff" value="Auto-Sync Off" class="btn btn-primary mb-2"> {{else}}
        <input style="width:40%" type="submit" name="syncon" value="Auto-Sync On" class="btn btn-primary mb-2"> {{end}}
    </form>
    <script src="http://apps.bdimg.com/libs/jquery/2.1.4/jquery.min.js"></script>
    <script>
        $(function() {
            var ws = new WebSocket('ws://' + window.location.host + '/host/ws')
            ws.onmessage = function(event) {
                // $('<li>').text(event.data).appendTo($ul);
                var data = JSON.parse(event.data);
                var CPU = data.CPU;
                var Memory = data.Memory;
                var HostName = data.HostName;
                var OS = data.OS;
                var Disk = data.DiskSpace;
                hostname.text("Hostname: " + HostName);
                hostos.text("OS: " + OS);
                cpumodel.text("CPU Model: " + CPU.CPUModelandFrequency);
                totalmemory.text("Total Memory: " + Memory.TotalMemory + " KB");
                usedmemory.text("Used Memory: " + Memory.UsedMemory + " KB");
                memoryusage.text("Memory Usage: " + Memory.MemoryPercentage + " %");
                swap.text("Swap On or Off: " + Memory.SWAPonoff);
                totaldisk.text("TotalDisk: " + Disk.TotalDisk + " B");
                useddisk.text("Used Disk: " + Disk.UsedDisk + " B");
                availdisk.text("Availble Disk: " + Disk.AvaileDisk + " B");
                useddiskpercentage.text("Used Disk Percentage: " + Disk.DiskPercentage + " %");
                var percentagelist = "";
                for (let index = 0; index < CPU.CPUPercentage.length; index++) {
                    percentagelist = percentagelist + "<a href=\"#\" class=\"list-group-item list-group-item-action\">" + "CPU " + index + " " + CPU.CPUPercentage[index] + " %</a>";
                }
                cpupercentage.html(percentagelist);
            };
            // ws.onopen = function(event) {
            //     //var opendata = JSON.parse(event.data);
            //     alert("The Socket has opended");

            // };

            var hostname = $("#Hostname");
            var hostos = $("#HostOS");
            var cpumodel = $("#CPUModel");
            var totalmemory = $("#TotalMemory");
            var usedmemory = $("#UsedMemory");
            var memoryusage = $("#MemoryUsage");
            var swap = $("#SWAP");
            var totaldisk = $("#TotalDisk");
            var useddisk = $("#UsedDisk");
            var availdisk = $("#AvailbleDisk");
            var useddiskpercentage = $("#UsedDiskPercentage");
            var cpupercentage = $("#CPUpercentage");
            // alert(cpupercentage.html());
        });
    </script>
</div>
{{end}}