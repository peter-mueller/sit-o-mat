. ./environment.sh

echo "Create Users:"
echo "\`\`\`"
./sitomatadmin register -name p.mueller
./sitomatadmin register -name Renegade
echo "\`\`\`"

echo "# Create Workplaces:"
echo "\`\`\`"
./sitomatadmin addworkplace -location 5.14 -name Master-Throne
./sitomatadmin addworkplace -location 5.14 -name One-And-Only
./sitomatadmin addworkplace -location 5.14 -name Coffeemaker
./sitomatadmin addworkplace -location 5.14 -name Soundmaster

echo "\`\`\`"


echo "# Assign Users to Workplaces"
echo "\`\`\`"
./sitomatadmin assignworkplaces
echo "\`\`\`"
