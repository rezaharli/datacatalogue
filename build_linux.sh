export GOOS=linux

tanggal=`date +%Y%m%d%H%M%S`
foldername="$CI_PROJECT_NAME-$tanggal"
rm -r "builds"
mkdir "builds"
mkdir "builds/${foldername}"
mkdir "builds/${foldername}/webapp"
mkdir "builds/${foldername}/webapp/config"
mkdir "builds/${foldername}/webapp/views"

cd webapp

go build -o datacatalogue
mv datacatalogue "../builds/${foldername}/webapp"

cd views
cp -r dist/ "../../builds/${foldername}/webapp/views"

cd ..
cp config/app.json.template "../builds/${foldername}/webapp/config/"
cd ../builds/
zip -r "${foldername}.tar.gz" "${foldername}"
echo "${foldername}"
rm -r "${foldername}"

export GOOS=darwin
