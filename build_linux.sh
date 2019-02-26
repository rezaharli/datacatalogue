export GOOS=linux

foldername="$CI_PROJECT_NAME-$CI_COMMIT_TAG"
rm -r "builds"
mkdir builds
mkdir "builds/${foldername}"
mkdir "builds/${foldername}/webapp"
mkdir "builds/${foldername}/webapp/config"
mkdir "builds/${foldername}/webapp/views"

cd webapp
go build -o datacatalogue
mv datacatalogue "../builds/${foldername}/webapp"

cd views
yarn build
cp -r dist/ "../../builds/${foldername}/"

cd ..
cp config/app.json.template "../builds/${foldername}/webapp/config/"
cd ../builds/
tar -zcf "${foldername}.tar.gz" "${foldername}"
rm -r "${foldername}"

export GOOS=darwin
