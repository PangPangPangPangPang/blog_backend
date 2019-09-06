#/bin/bash
# . ./update_bundle.sh

if [ ! -d "$HOME/blog_backend_temp" ]; then
    cd $HOME
    git clone http://github.com/PangPangPangPangPang/blog_backend.git $HOME/blog_backend_temp
    echo $HOME
    echo ---not exist---
else
    cd $HOME/blog_backend_temp
    git pull
fi

rm -rf $HOME/blog_backend/resource
cp -rf $HOME/blog_backend_temp/resource $HOME/blog_backend/
echo -----------------------------------------------------
echo ----------update article success---------------------
echo -----------------------------------------------------


if [ ! -d "$HOME/blog_frontend" ]; then
    cd $HOME
    git clone http://github.com/PangPangPangPangPang/blog_frontend.git
else
    cd $HOME/blog_frontend
    git pull
fi

rm -rf $HOME/blog_backend/static/*
cp -r $HOME/blog_frontend/build/* $HOME/blog_backend/static/
cp -r $HOME/blog_backend_temp/img/* $HOME/blog_backend/static/

echo -----------------------------------------------------
echo ----------update frontend bundle success-------------
echo -----------------------------------------------------
