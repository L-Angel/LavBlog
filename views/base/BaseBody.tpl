<!DOCTYPE html>
<html lang="en">
<head>
{{template "base/Header.tpl" .}}
</head>
<body>
{{template "base/NavLeft.tpl" .}}
{{template "base/Head.tpl" .}}
<main>
<div class="lav-body container">
    <div id="content-wrapper">
        <div class="mui--appbar-height"></div>
        <div class="mui-container-fluid">
            {{template "body" .}}
        </div>
    </div>
</div>
</main>
{{template "base/Footer.tpl" .}}
</body>
</html>