. ./environment.sh

echo "Create Users:"
echo "\`\`\`"
./sitomatadmin register -name p.mueller
./sitomatadmin register -name rene.zarwel
./sitomatadmin register -name franz.test
echo "\`\`\`"

echo "# Create Workplaces:"
echo "\`\`\`"
./sitomatadmin addworkplace -location A5.12 -name uhu
./sitomatadmin addworkplace -location A5.12 -name bhu
./sitomatadmin addworkplace -location A5.12 -name chu
./sitomatadmin addworkplace -location A5.12 -name dhu

./sitomatadmin addworkplace -location A5.13 -name asdf
./sitomatadmin addworkplace -location A5.14 -name wef
./sitomatadmin addworkplace -location A5.1 -name ehr
./sitomatadmin addworkplace -location A5.10 -name dwerhhu
echo "\`\`\`"


echo "# Assign Users to Workplaces"
echo "\`\`\`"
./sitomatadmin assignworkplaces
echo "\`\`\`"
