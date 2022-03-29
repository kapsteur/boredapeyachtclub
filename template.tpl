<html>
    <head>
        <meta http-equiv="refresh" content="{{.Refresh}};URL=/random">
        <style>
            body {
                background-color: rgba({{.R}},{{.G}},{{.B}},{{.A}});
            }

            img {
                max-width: 1000px;
                max-height: 1000px;
                width: auto;
                height: auto;
                position: fixed;
                left: 50%;
                bottom: 0px;
                transform: translate(-50%, 0);
                margin: 0 auto;
            }
        </style>
    </head>
    <body>
        <img src="{{.Img}}"/>
    </body>
</html>