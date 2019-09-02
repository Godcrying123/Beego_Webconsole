{{define "navbar"}}
<header>
    <div>
        <nav class="navbar navbar-expand-lg navbar-dark bg-primary">
            <a class="navbar-brand" href="/">SMA Web-Console</a>
            <div class="collapse navbar-collapse" id="navbarColor02">
                <ul class="navbar-nav mr-auto">
                    <li class="nav-item">
                        <a class="nav-link" href="{{.sshUrl}}file/" target="_blank">Files<span class="sr-only">(current)</span></a>
                    </li>
                    <li class="nav-item">
                        <a class="nav-link" href="/service/" target="_blank">Services Edit<span class="sr-only">(current)</span></a>
                    </li>
                    <li class="nav-item">
                        <a class="nav-link" href="/step/" target="_blank">Steps Edit<span class="sr-only">(current)</span></a>
                    </li>
                    <li class="nav-item">
                        <a class="nav-link" href="/machine/">Machine<span class="sr-only">(current)</span></a>
                    </li>
                    <li class="nav-item">
                        <a class="nav-link" href="/file/">Xlooklook<span class="sr-only">(current)</span></a>
                    </li>
                    <li class="nav-item dropdown">
                        <a class="nav-link dropdown-toggle" href="#" id="navbarDropdownhost" role="button" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">Host Connect</a>
                        <div class="dropdown-menu" aria-labelledby="navbarDropdownhost">
                            {{range $key,$value := .machine}}
                            <a class="dropdown-item" href="/node/{{$key}}" , target="_blank">{{$key}}->{{$value.HostName}}</a> {{end}}
                        </div>
                    </li>
                </ul>
                <div class="form-inline my-2 my-lg-0">
                    <ul class="navbar-nav mr-auto">
                        <li class="nav-item dropdown">
                            <a class="nav-link dropdown-toggle" href="#" id="navbarDropdownMenuLink1" data-toggle="modal" data-target="#exampleModal" aria-haspopup="true" aria-expanded="false">All Steps</a>
                            <div class="modal fade" id="exampleModal" tabindex="-1" role="dialog" aria-labelledby="exampleModalLabel" aria-hidden="true">
                                <div class="modal-dialog" role="document">
                                    {{template "step" .}}
                                </div>
                            </div>
                        </li>
                        <li class="nav-item dropdown">
                            <a class="nav-link dropdown-toggle" href="#" id="navbarDropdownMenuLink3" data-toggle="modal" data-target="#exampleModal1" aria-haspopup="true" aria-expanded="false" data-toggle=".bd-example-modal-lg">All Services</a>
                            <div class="modal fade" id="exampleModal1" tabindex="-1" role="dialog" aria-labelledby="exampleModalLabel" aria-hidden="true">
                                <div class="modal-dialog" role="document">
                                    {{template "service" .}}
                                </div>
                            </div>
                        </li>
                        <li class="nav-item dropdown">
                            <a class="nav-link dropdown-toggle" href="#" id="navbarDropdownMenuLink2" data-toggle="modal" data-target=".bd-example-modal-lg" aria-haspopup="true" aria-expanded="false">Host Info</a>
                            <div class="modal fade bd-example-modal-lg" tabindex="-1" role="dialog" aria-labelledby="myLargeModalLabel" aria-hidden="true">
                                <div class="modal-dialog modal-lg">
                                    <div class="modal-content">
                                        {{template "host" .}}
                                    </div>
                                </div>
                            </div>
                        </li>
                    </ul>
                </div>
            </div>
        </nav>
    </div>
</header>
{{end}}