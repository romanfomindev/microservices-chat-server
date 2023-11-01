package repositories

//go:generate sh -c "rm -rf mocks && mkdir -p mocks"
//go:generate ../../bin/minimock -i Chat -o ./mocks/ -s "_minimock.go"
//go:generate ../../bin/minimock -i ChatUser -o ./mocks/ -s "_minimock.go"
