{{define "service"}}
<!-- <div style="width:20%; float: right;"> -->
<div class="modal-content">
    <div class="modal-header">
        <h5 class="modal-title" id="exampleModalLabel">Services</h5>
        <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
    </div>
    <div class="modal-body">
        <table class="table table-responsive table-hover table-striped">
            <thead class="thead-dark">
                <tr>
                    <th colspan="3">SerVice List</th>
                </tr>
            </thead>
            {{range .services}}
            <tbody>
                <tr class="clickable" data-toggle="collapse" data-target="#{{.ServiceName}}" aria-expanded="false" aria-controls="{{.ServiceName}}">
                    <td colspan="3">{{.ServiceName}}.Service</td>
                </tr>
            </tbody>
            <tbody id="{{.ServiceName}}" class="collapse">
                <tr>
                    <td>{{.ServiceVersion}}</td>
                    <td>{{.ActiveStatus}}</td>
                    <td>{{.RunningStatus}}</td>
                </tr>
                <tr>
                    <td style="overflow:hidden" colspan="3">{{.ServiceStatus}}</td>
                </tr>
            </tbody>
            {{end}}
        </table>
        <div class="modal-footer">
            <form method="POST" action="/" enctype="multipart/form-data" target="iframe">
                <input type="file" name="importfile" />
                <input type="button" class="btn btn-primary" value="Edit All" onclick="javascript:window.location.href='/service/';" />
                <input type="submit" class="btn btn-primary" name="importall" value="Import All" />
            </form>
        </div>
        <script src="http://apps.bdimg.com/libs/jquery/2.1.4/jquery.min.js"></script>
        <script>
            $(function() {
                var ws = new WebSocket('ws://' + window.location.host + '/service/ws')
                ws.onmessage = function(event) {
                    // $('<li>').text(event.data).appendTo($ul);
                    var data = JSON.parse(event.data);
                    var active = data.ActiveStatus;
                    var running = data.RunningStatus
                    if (active != "active" || running != "running") {
                        console.log("The Service " + data.ServiceName + " Status is Not Good");

                    };
                    // console.log(data.ServiceName);
                    // console.log(data.ActiveStatus);
                };
                ws.onopen = function() {
                    console.log("WebSocket Data Has Started!")
                };
            });
        </script>
    </div>
</div>
{{end}}