<section class="content-header">
    <h1>
        修改用户
    </h1>
    <ol class="breadcrumb">
        <li><a href="#"><i class="fa fa-dashboard"></i> 用户管理</a></li>
        <li><a href="#">创建用户</a></li>
    </ol>
</section>

<section class="content">
    <div class="row">
        <div class="col-md-12">
            <div class="box box-primary">

                <form role="form" action="/admin/user/add" method="post" enctype="multipart/form-data">
                    <div class="box-body">
                        <input type="hidden" name="user_id" value="{{ .userInfo.Id }}">
                        <div class="form-group">
                            <label for="exampleInputEmail1">用户名</label>
                            <input type="text" class="form-control" id="" name="username" value="{{ .userInfo.Username }}">
                        </div>
                        <div class="form-group">
                            <label for="exampleInputPassword1">密码</label>
                            <input type="password" class="form-control" id="exampleInputPassword1" name="password" value="{{ .userInfo.Password }}">
                        </div>
                        <div class="form-group">
                            <label for="exampleInputFile">上传头像</label>
                            <input type="file" id="exampleInputFile" name="image">

                        </div>
                        <div class="form-group">
                            <img src="{{ .userInfo.Img }}" alt="">
                        </div>

                    </div>

                    <div class="box-footer">
                        <button type="submit" class="btn btn-primary">提交</button>
                    </div>

                </form>
            </div>

        </div>

    </div>
</section>