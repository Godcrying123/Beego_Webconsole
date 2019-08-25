{{define "navbar"}}
<header>
    <div>
        <nav class="navbar navbar-expand-lg navbar-dark bg-primary">
            <a class="navbar-brand" href="/">SMA Web-Console</a>
            <div class="collapse navbar-collapse" id="navbarColor02">
                <ul class="navbar-nav mr-auto">
                    <li class="nav-item">
                        <a class="nav-link" href="/file/">Files<span class="sr-only">(current)</span></a>
                    </li>
                    <li class="nav-item">
                        <a class="nav-link" href="/service/">Services Edit<span class="sr-only">(current)</span></a>
                    </li>
                    <li class="nav-item">
                        <a class="nav-link" href="/step/">Steps Edit<span class="sr-only">(current)</span></a>
                    </li>
                    <li class="nav-item">
                        <a class="nav-link" href="/file/">Xlooklook<span class="sr-only">(current)</span></a>
                    </li>
                </ul>
                <form class="form-inline my-2 my-lg-0">
                    <ul class="navbar-nav mr-auto">
                        <li class="nav-item dropdown">
                            <a class="nav-link dropdown-toggle" href="#" id="navbarDropdownMenuLink1" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">All Steps</a>
                            <div class="dropdown-menu" aria-labelledby="navbarDropdownMenuLink1">
                                {{template "step" .}}
                            </div>
                        </li>
                        <li class="nav-item dropdown">
                            <a class="nav-link dropdown-toggle" href="#" id="navbarDropdownMenuLink2" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">Host Info</a>
                            <div style="width: 200%;" class="dropdown-menu" aria-labelledby="navbarDropdownMenuLink2">
                                {{template "host" .}}
                            </div>
                        </li>
                        <li class="nav-item dropdown">
                            <a class="nav-link dropdown-toggle" href="#" id="navbarDropdownMenuLink3" data-toggle="modal" aria-haspopup="true" aria-expanded="false" data-toggle=".bd-example-modal-lg">All Services</a>
                            <div class="dropdown-menu" aria-labelledby="navbarDropdownMenuLink3">
                                {{template "service" .}}
                            </div>
                        </li>
                    </ul>
                </form>
            </div>
        </nav>
    </div>
</header>
{{end}}