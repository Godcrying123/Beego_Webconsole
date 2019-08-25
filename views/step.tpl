{{define "step"}}
<!-- <div style="float: left; width: 22%; display: inline;"> -->
<div class="modal-content">
    <div class="modal-header">
        <h5 class="modal-title" id="exampleModalLabel">Steps</h5>
        <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
    </div>
    <div id="accordion" class="modal-body">
        {{range .stepsData}}
        <div class="card ">
            <div class="card-header" id="headingOne">
                <h5 class="mb-0">
                    <button class="btn btn-link" data-toggle="collapse" data-target="#collapseOne" aria-expanded="true" aria-controls="collapseOne">{{.StepTitle}}</button>
                </h5>
            </div>
            <div id="collapseOne" class="collapse show" aria-labelledby="headingOne" data-parent="#accordion">
                <div>
                    {{range .SubSteps}}
                    <a href="#" class="list-group-item list-group-item-action">{{.StepName}}</a>
                    <input type="hidden" name="" value="{{.StepSummary}}">
                    <input type="hidden" name="" value="{{.StepCommand}}"> {{end}}
                </div>
            </div>
        </div>
        {{end}}
    </div>
    <div class="modal-footer">
        <form action="/" method="POST" enctype="multipart/form-data">
            <div>
                <input type="file" name="importfilestep" />
                <input type="button" class="btn btn-primary" value="Edit" onclick="javascript:window.location.href='/step/';"></input>
                <input type="submit" class="btn btn-primary" name="importallsteps" value="Import" />
            </div>
        </form>
    </div>
</div>
<!-- <div>
    <div id="accordion">
        {{range .stepsData}}
        <div class="card ">
            <div class="card-header" id="headingOne">
                <h5 class="mb-0">
                    <button class="btn btn-link" data-toggle="collapse" data-target="#{{.StepTitle}}" aria-expanded="true" aria-controls="{{.StepTitle}}">{{.StepTitle}}</button>
                </h5>
            </div>
            <div id="{{.StepTitle}}" class="collapse show" aria-labelledby="headingOne" data-parent="#accordion">
                <div>
                    {{range .SubSteps}}
                    <a href="#" class="list-group-item list-group-item-action">{{.StepName}}</a>
                    <input type="hidden" name="" value="{{.StepSummary}}">
                    <input type="hidden" name="" value="{{.StepCommand}}"> {{end}}
                </div>
            </div>
        </div>
        {{end}}
    </div>
    <form action="/" method="POST" enctype="multipart/form-data">
        <div>
            <input type="file" name="importfilestep" />
            <input type="button" class="btn btn-primary" value="Edit" onclick="javascript:window.location.href='/step/';"></input>
            <input type="submit" class="btn btn-primary" name="importallsteps" value="Import" />
        </div>
    </form>
</div> -->
{{end}}