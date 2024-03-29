{{define "task"}}
<div class="modal-content">
    <div class="modal-header">
        <h5 class="modal-title" id="exampleModalLabel">Tasks</h5>
        <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
    </div>
    <div class="modal-body">
        <div class="container">
            <ul>
                <h3>This is the auto task executor</h3>
                <li>You can click the main test run button to execute all sub tasks</li>
                <li>You can click the sub task run button to execute this specific task</li>
                <li>You can import the json file to get the task list</li>
                <li>You can customize each step but cannot customize the main step</li>
            </ul>
        </div>
        <hr>
        <div>
            <form method="POST" action="/" enctype="multipart/form-data">
                <!-- target="iframetaskrun"  -->
                <div class="accordion container" id="accordionExample">
                    {{range $key, $value := .taskData}}
                    <form action="/" method="POST" enctype="multipart/form-data">
                        <div class="card">
                            <div class="card-header" id="headingOne">
                                <div class="mb-0">
                                    <button class="btn btn-link" type="button" data-toggle="collapse" data-target="#task{{$value.ID}}" aria-expanded="true" aria-controls="collapseOne">{{$key}}</button>
                                    <input type="submit" name="runAllTask" value="run all tasks"></input>
                                    <input type="hidden" name="AllTaskDetails" value="{{$key}}">
                                </div>
                            </div>

                            <div id="task{{$value.ID}}" class="collapse show" aria-labelledby="headingOne" data-parent="#accordionExample">
                                <div class="card-body">
                                    <table class="table">
                                        <form action="/task" method="POST" enctype="multipart/form-data"></form>
                                        {{range $value.SubTasks}}
                                        <form action="/" method="POST" enctype="multipart/form-data">
                                            <tr>
                                                <td><input type="text" value="{{.TaskSummary}}"></td>
                                                <td><input type="text" name="TaskNode" value="{{.TaskNode}}"></td>
                                                <td><input type="text" name="TaskCommand" value="{{.TaskCommand}}"></td>
                                                <td><input type="submit" value="Run it!" name="runTask"></td>
                                            </tr>
                                        </form>
                                        {{end}}
                                    </table>
                                </div>
                            </div>
                        </div>
                    </form>
                    {{end}}
                    <div style="text-align: right;">
                        <input type="file" name="importfiletasks" />
                        <input type="submit" class="btn btn-primary" name="importalltasks" value="Import All Tasks" />
                    </div>
                </div>
            </form>
            <iframe id="iframetaskrun" name="iframetaskrun" style="display:none;"></iframe>
        </div>
    </div>
</div>

{{end}}