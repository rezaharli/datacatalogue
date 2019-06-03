export GOOS=linux

tanggal=`date +%Y%m%d_%H%M`
foldername="$CI_PROJECT_NAME-$tanggal"
rm -r "builds"
mkdir "builds"
mkdir "builds/${foldername}"
mkdir "builds/${foldername}/webapp"
mkdir "builds/${foldername}/webapp/config"
mkdir "builds/${foldername}/webapp/views"
mkdir "builds/${foldername}/webapp/queryfiles"

cd webapp

go build -o datacatalogue
mv datacatalogue "../builds/${foldername}/webapp"

cp -r views/dist/ "../builds/${foldername}/webapp/views"

cp config/app.json.template "../builds/${foldername}/webapp/config/"

cp -a "queryfiles/." "../builds/${foldername}/webapp/queryfiles/"

cd ../builds/
zip -r "${foldername}.zip" "${foldername}"
echo ""
echo "Build:"
echo "http://go.eaciit.com/files/datacatalogue/${foldername}.zip"
echo ""
rm -r "${foldername}"

export GOOS=darwin
