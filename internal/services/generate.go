package services

//go:generate sh -c "rm -rf mocks && mkdir -p mocks"
//go:generate ../../bin/minimock -i ChatService -o ./mocks/ -s "_minimock.go"
//go:generate ../../bin/minimock -i Implementation -o ./mocks/ -s "_minimock.go"



