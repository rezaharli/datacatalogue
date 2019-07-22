export GOOS=linux

tanggal=`date +%Y%m%d_%H%M`
foldername="$CI_PROJECT_NAME-$tanggal"
rm -r "builds"
mkdir "builds"
mkdir "builds/${foldername}"
mkdir "builds/${foldername}/build/webapp"
mkdir "builds/${foldername}/build/webapp/config"
mkdir "builds/${foldername}/build/webapp/views"
mkdir "builds/${foldername}/build/webapp/queryfiles"

cd webapp

go build -o datacatalogue
mv datacatalogue "../builds/${foldername}/build/webapp"

cp -r views/dist/ "../builds/${foldername}/build/webapp/views"

cp config/app.json.template "../builds/${foldername}/build/webapp/config/"

cp -a "queryfiles/." "../builds/${foldername}/build/webapp/queryfiles/"

cd ../builds/
zip -r "${foldername}.zip" "${foldername}"
echo ""
echo "Build:"
echo "http://go.eaciit.com/files/datacatalogue/${foldername}.zip"
echo ""
rm -r "${foldername}"

export GOOS=darwin
