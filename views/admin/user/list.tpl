<!-- Content Header (Page header) -->

<section class="content-header">
    <h1>
        用户列表
        {{/*<small>preview of simple tables</small>*/}}
    </h1>
    <ol class="breadcrumb">
        <li><a href="#"><i class="fa fa-dashboard"></i> 用户管理</a></li>
        <li><a href="#">用户列表</a></li>
        {{/*<li class="active">Simple</li>*/}}
    </ol>
</section>

<!-- Main content -->
<section class="content">
    <div class="row">
        <div class="col-md-12">
            <div class="box">
                {{/*<div class="box-header with-border">*/}}
                    {{/*<h3 class="box-title">Bordered Table</h3>*/}}
                {{/*</div>*/}}
                <!-- /.box-header -->
                <div class="box-body">
                    <table class="table table-bordered">
                        <tr>
                            <th>用户Id</th>
                            <th>用户名</th>
                            <th>添加时间</th>
                            <th>操作</th>
                        </tr>
                        {{range $index, $item := .userList}}
                        <tr>
                            <td>{{ $item.id }}</td>
                            <td>
                                {{ $item.username }}
                            </td>
                            <td>{{ $item.create_time }}</td>
                            <td><a href="/admin/user/edit/{{ $item.id }}" class="btn btn-info">编辑</a>
                            <a class="btn btn-danger">删除</a></td>
                        </tr>
                        {{end}}
                    </table>
                </div>
                <!-- /.box-body -->
                <div class="box-footer clearfix">
                    <ul class="pagination">

                    </ul>
                </div>
            </div>
        </div>

    </div>

</section>


<!-- /.content -->