package main

import (
	"reflect"
	"testing"
)

func TestAdaptee_ExistingMethod(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Adaptee{}
		})
	}
}

func TestAdapter_ExpectedMethod(t *testing.T) {
	type fields struct {
		adaptee *Adaptee
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Adapter{
				adaptee: tt.fields.adaptee,
			}
		})
	}
}

func TestCaretaker(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestChainedReceiver_Finish(t *testing.T) {
	type fields struct {
		canHandle string
		next      *ChainedReceiver
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ChainedReceiver{
				canHandle: tt.fields.canHandle,
				next:      tt.fields.next,
			}
			if err := r.Finish(); (err != nil) != tt.wantErr {
				t.Errorf("Finish() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestChainedReceiver_Handle(t *testing.T) {
	type fields struct {
		canHandle string
		next      *ChainedReceiver
	}
	type args struct {
		what string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ChainedReceiver{
				canHandle: tt.fields.canHandle,
				next:      tt.fields.next,
			}
			if err := r.Handle(tt.args.what); (err != nil) != tt.wantErr {
				t.Errorf("Handle() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestChainedReceiver_SetNext(t *testing.T) {
	type fields struct {
		canHandle string
		next      *ChainedReceiver
	}
	type args struct {
		next *ChainedReceiver
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ChainedReceiver{
				canHandle: tt.fields.canHandle,
				next:      tt.fields.next,
			}
		})
	}
}

func TestColleague1_GetState(t *testing.T) {
	type fields struct {
		mediator Mediator
		state    string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Colleague1{
				mediator: tt.fields.mediator,
				state:    tt.fields.state,
			}
			if got := c.GetState(); got != tt.want {
				t.Errorf("GetState() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestColleague1_SetMediator(t *testing.T) {
	type fields struct {
		mediator Mediator
		state    string
	}
	type args struct {
		mediator Mediator
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Colleague1{
				mediator: tt.fields.mediator,
				state:    tt.fields.state,
			}
		})
	}
}

func TestColleague1_SetState(t *testing.T) {
	type fields struct {
		mediator Mediator
		state    string
	}
	type args struct {
		state string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Colleague1{
				mediator: tt.fields.mediator,
				state:    tt.fields.state,
			}
		})
	}
}

func TestColleague2_GetState(t *testing.T) {
	type fields struct {
		mediator Mediator
		state    int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Colleague2{
				mediator: tt.fields.mediator,
				state:    tt.fields.state,
			}
			if got := c.GetState(); got != tt.want {
				t.Errorf("GetState() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestColleague2_SetMediator(t *testing.T) {
	type fields struct {
		mediator Mediator
		state    int
	}
	type args struct {
		mediator Mediator
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Colleague2{
				mediator: tt.fields.mediator,
				state:    tt.fields.state,
			}
		})
	}
}

func TestColleague2_SetState(t *testing.T) {
	type fields struct {
		mediator Mediator
		state    int
	}
	type args struct {
		state int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Colleague2{
				mediator: tt.fields.mediator,
				state:    tt.fields.state,
			}
		})
	}
}

func TestComposite_AddChild(t *testing.T) {
	type fields struct {
		children []InterfaceX
	}
	type args struct {
		child InterfaceX
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Composite{
				children: tt.fields.children,
			}
		})
	}
}

func TestComposite_MethodA(t *testing.T) {
	type fields struct {
		children []InterfaceX
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Composite{
				children: tt.fields.children,
			}
		})
	}
}

func TestConcreteMediator_SetColleagueC1(t *testing.T) {
	type fields struct {
		c1 Colleague1
		c2 Colleague2
	}
	type args struct {
		c1 Colleague1
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ConcreteMediator{
				c1: tt.fields.c1,
				c2: tt.fields.c2,
			}
		})
	}
}

func TestConcreteMediator_SetColleagueC2(t *testing.T) {
	type fields struct {
		c1 Colleague1
		c2 Colleague2
	}
	type args struct {
		c2 Colleague2
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ConcreteMediator{
				c1: tt.fields.c1,
				c2: tt.fields.c2,
			}
		})
	}
}

func TestConcreteMediator_SetState(t *testing.T) {
	type fields struct {
		c1 Colleague1
		c2 Colleague2
	}
	type args struct {
		s string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ConcreteMediator{
				c1: tt.fields.c1,
				c2: tt.fields.c2,
			}
		})
	}
}

func TestConcreteNodeX_Accept(t *testing.T) {
	type args struct {
		visitor Visitor
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := ConcreteNodeX{}
		})
	}
}

func TestConcreteNodeY_Accept(t *testing.T) {
	type args struct {
		visitor Visitor
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := ConcreteNodeY{}
		})
	}
}

func TestConcreteObserverA_SetModel(t *testing.T) {
	type fields struct {
		model     *Subject
		viewState string
	}
	type args struct {
		s *Subject
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ca := &ConcreteObserverA{
				model:     tt.fields.model,
				viewState: tt.fields.viewState,
			}
		})
	}
}

func TestConcreteObserverA_Update(t *testing.T) {
	type fields struct {
		model     *Subject
		viewState string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ca := &ConcreteObserverA{
				model:     tt.fields.model,
				viewState: tt.fields.viewState,
			}
		})
	}
}

func TestConcreteReportA_Execute(t *testing.T) {
	type fields struct {
		receiver *Receiver
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ConcreteReportA{
				receiver: tt.fields.receiver,
			}
		})
	}
}

func TestConcreteReportB_Execute(t *testing.T) {
	type fields struct {
		receiver *Receiver
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ConcreteReportB{
				receiver: tt.fields.receiver,
			}
		})
	}
}

func TestConcreteVisitor_Visit(t *testing.T) {
	type args struct {
		node Node
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := ConcreteVisitor{}
		})
	}
}

func TestContext_Op1(t *testing.T) {
	type fields struct {
		state State
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Context{
				state: tt.fields.state,
			}
		})
	}
}

func TestContext_Op2(t *testing.T) {
	type fields struct {
		state State
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Context{
				state: tt.fields.state,
			}
		})
	}
}

func TestContext_SetState(t *testing.T) {
	type fields struct {
		state State
	}
	type args struct {
		state State
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Context{
				state: tt.fields.state,
			}
		})
	}
}

func TestFastAlgo_FindBreadth(t *testing.T) {
	type args struct {
		set []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &FastAlgo{}
			if got := n.FindBreadth(tt.args.set); got != tt.want {
				t.Errorf("FindBreadth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlightFactory_CreateInvoice(t *testing.T) {
	tests := []struct {
		name string
		want Invoice
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := FlightFactory{}
			if got := f.CreateInvoice(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateInvoice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlightFactory_CreateReservation(t *testing.T) {
	tests := []struct {
		name string
		want Reservation
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := FlightFactory{}
			if got := f.CreateReservation(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateReservation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlightReservationImpl_AddExtraLuggageAllowance(t *testing.T) {
	type fields struct {
		reservationDate string
		luggageAllowed  int
	}
	type args struct {
		peices int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := FlightReservationImpl{
				reservationDate: tt.fields.reservationDate,
				luggageAllowed:  tt.fields.luggageAllowed,
			}
		})
	}
}

func TestFlightReservationImpl_CalculateCancellationFee(t *testing.T) {
	type fields struct {
		reservationDate string
		luggageAllowed  int
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := FlightReservationImpl{
				reservationDate: tt.fields.reservationDate,
				luggageAllowed:  tt.fields.luggageAllowed,
			}
			if got := r.CalculateCancellationFee(); got != tt.want {
				t.Errorf("CalculateCancellationFee() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlightReservationImpl_GetReservationDate(t *testing.T) {
	type fields struct {
		reservationDate string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := FlightReservationImpl{
				reservationDate: tt.fields.reservationDate,
			}
			if got := r.GetReservationDate(); got != tt.want {
				t.Errorf("GetReservationDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlightReservationImpl_GetReservationDate1(t *testing.T) {
	type fields struct {
		reservationDate string
		luggageAllowed  int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := FlightReservationImpl{
				reservationDate: tt.fields.reservationDate,
				luggageAllowed:  tt.fields.luggageAllowed,
			}
			if got := r.GetReservationDate(); got != tt.want {
				t.Errorf("GetReservationDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlightReservationImpl_SetReservationDate(t *testing.T) {
	type fields struct {
		reservationDate string
	}
	type args struct {
		date string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := FlightReservationImpl{
				reservationDate: tt.fields.reservationDate,
			}
		})
	}
}

func TestGetFactory(t *testing.T) {
	type args struct {
		vertical string
	}
	tests := []struct {
		name string
		args args
		want AbstractFactory
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetFactory(tt.args.vertical); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFactory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetMyClass(t *testing.T) {
	tests := []struct {
		name string
		want *MyClass
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMyClass(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMyClass() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHotelBoutiqueProxy_Book(t *testing.T) {
	type fields struct {
		subject *HotelBoutique
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &HotelBoutiqueProxy{
				subject: tt.fields.subject,
			}
		})
	}
}

func TestHotelBoutique_Book(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &HotelBoutique{}
		})
	}
}

func TestHotelFactory_CreateInvoice(t *testing.T) {
	tests := []struct {
		name string
		want Invoice
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := HotelFactory{}
			if got := f.CreateInvoice(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateInvoice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHotelFactory_CreateReservation(t *testing.T) {
	tests := []struct {
		name string
		want Reservation
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := HotelFactory{}
			if got := f.CreateReservation(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateReservation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHotelReservationImpl_CalculateCancellationFee(t *testing.T) {
	type fields struct {
		reservationDate string
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := HotelReservationImpl{
				reservationDate: tt.fields.reservationDate,
			}
			if got := r.CalculateCancellationFee(); got != tt.want {
				t.Errorf("CalculateCancellationFee() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHotelReservationImpl_GetReservationDate(t *testing.T) {
	type fields struct {
		reservationDate string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := HotelReservationImpl{
				reservationDate: tt.fields.reservationDate,
			}
			if got := r.GetReservationDate(); got != tt.want {
				t.Errorf("GetReservationDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHotelReservationImpl_GetReservationDate1(t *testing.T) {
	type fields struct {
		reservationDate string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := HotelReservationImpl{
				reservationDate: tt.fields.reservationDate,
			}
			if got := r.GetReservationDate(); got != tt.want {
				t.Errorf("GetReservationDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHotelReservationImpl_SetReservationDate(t *testing.T) {
	type fields struct {
		reservationDate string
	}
	type args struct {
		date string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := HotelReservationImpl{
				reservationDate: tt.fields.reservationDate,
			}
		})
	}
}

func TestInstitutionSeller_CancelReservation(t *testing.T) {
	type args struct {
		charge float64
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := InstitutionSeller{}
		})
	}
}

func TestInvoker_Run(t *testing.T) {
	type fields struct {
		repository []Report
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Invoker{
				repository: tt.fields.repository,
			}
		})
	}
}

func TestInvoker_Schedule(t *testing.T) {
	type fields struct {
		repository []Report
	}
	type args struct {
		cmd Report
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Invoker{
				repository: tt.fields.repository,
			}
		})
	}
}

func TestMasterAlgorithm_TemplateMethod(t *testing.T) {
	type fields struct {
		template Template
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &MasterAlgorithm{
				template: tt.fields.template,
			}
		})
	}
}

func TestMemento_GetState(t *testing.T) {
	type fields struct {
		serializedState string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Memento{
				serializedState: tt.fields.serializedState,
			}
			if got := m.GetState(); got != tt.want {
				t.Errorf("GetState() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyClass_GetAttrib(t *testing.T) {
	type fields struct {
		attrib string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &MyClass{
				attrib: tt.fields.attrib,
			}
			if got := c.GetAttrib(); got != tt.want {
				t.Errorf("GetAttrib() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyClass_SetAttrib(t *testing.T) {
	type fields struct {
		attrib string
	}
	type args struct {
		val string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &MyClass{
				attrib: tt.fields.attrib,
			}
		})
	}
}

func TestNaiveAlgo_FindBreadth(t *testing.T) {
	type args struct {
		set []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &NaiveAlgo{}
			if got := n.FindBreadth(tt.args.set); got != tt.want {
				t.Errorf("FindBreadth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewAdapter(t *testing.T) {
	tests := []struct {
		name string
		want *Adapter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAdapter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAdapter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewContext(t *testing.T) {
	tests := []struct {
		name string
		want *Context
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewContext(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewReservation(t *testing.T) {
	type args struct {
		vertical        string
		reservationDate string
	}
	tests := []struct {
		name string
		args args
		want Reservation
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewReservation(tt.args.vertical, tt.args.reservationDate); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewReservation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewReservationBuilder(t *testing.T) {
	tests := []struct {
		name string
		want ReservationBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewReservationBuilder(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewReservationBuilder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOriginator_GetMemento(t *testing.T) {
	type fields struct {
		state string
	}
	tests := []struct {
		name   string
		fields fields
		want   Memento
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Originator{
				state: tt.fields.state,
			}
			if got := o.GetMemento(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMemento() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOriginator_GetState(t *testing.T) {
	type fields struct {
		state string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Originator{
				state: tt.fields.state,
			}
			if got := o.GetState(); got != tt.want {
				t.Errorf("GetState() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOriginator_Restore(t *testing.T) {
	type fields struct {
		state string
	}
	type args struct {
		memento Memento
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Originator{
				state: tt.fields.state,
			}
		})
	}
}

func TestOriginator_SetState(t *testing.T) {
	type fields struct {
		state string
	}
	type args struct {
		state string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Originator{
				state: tt.fields.state,
			}
		})
	}
}

func TestPremiumReservation_Cancel(t *testing.T) {
	type fields struct {
		Reservation Reservation
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := PremiumReservation{
				Reservation: tt.fields.Reservation,
			}
		})
	}
}

func TestProfileDecorator(t *testing.T) {
	type args struct {
		fn Function
	}
	tests := []struct {
		name string
		args args
		want Function
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ProfileDecorator(tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProfileDecorator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReceiver_Action(t *testing.T) {
	type args struct {
		msg string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Receiver{}
		})
	}
}

func TestReservation_Cancel(t *testing.T) {
	type fields struct {
		sellerRef Seller
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Reservation{
				sellerRef: tt.fields.sellerRef,
			}
		})
	}
}

func TestSmallScaleSeller_CancelReservation(t *testing.T) {
	type args struct {
		charge float64
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := SmallScaleSeller{}
		})
	}
}

func TestSquareRoot(t *testing.T) {
	type args struct {
		n float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SquareRoot(tt.args.n); got != tt.want {
				t.Errorf("SquareRoot() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStateA_Op1(t *testing.T) {
	type args struct {
		c *Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &StateA{}
		})
	}
}

func TestStateA_Op2(t *testing.T) {
	type args struct {
		c *Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &StateA{}
		})
	}
}

func TestStateB_Op1(t *testing.T) {
	type args struct {
		c *Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &StateB{}
		})
	}
}

func TestStateB_Op2(t *testing.T) {
	type args struct {
		c *Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &StateB{}
		})
	}
}

func TestSubject_Attach(t *testing.T) {
	type fields struct {
		observers []Observer
		state     string
	}
	type args struct {
		observer Observer
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Subject{
				observers: tt.fields.observers,
				state:     tt.fields.state,
			}
		})
	}
}

func TestSubject_GetState(t *testing.T) {
	type fields struct {
		observers []Observer
		state     string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Subject{
				observers: tt.fields.observers,
				state:     tt.fields.state,
			}
			if got := s.GetState(); got != tt.want {
				t.Errorf("GetState() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubject_SetState(t *testing.T) {
	type fields struct {
		observers []Observer
		state     string
	}
	type args struct {
		newState string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Subject{
				observers: tt.fields.observers,
				state:     tt.fields.state,
			}
		})
	}
}

func TestTrip_AddReservation(t1 *testing.T) {
	type fields struct {
		reservations []Reservation
	}
	type args struct {
		r Reservation
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Trip{
				reservations: tt.fields.reservations,
			}
		})
	}
}

func TestTrip_CalculateCancellationFee(t1 *testing.T) {
	type fields struct {
		reservations []Reservation
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Trip{
				reservations: tt.fields.reservations,
			}
			if got := t.CalculateCancellationFee(); got != tt.want {
				t1.Errorf("CalculateCancellationFee() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVariantA_Step1(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &VariantA{}
		})
	}
}

func TestVariantA_Step2(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &VariantA{}
		})
	}
}

func Test_client(t *testing.T) {
	type args struct {
		s Strategy
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := client(tt.args.s); got != tt.want {
				t.Errorf("client() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reservationBuilder_Build(t *testing.T) {
	type fields struct {
		vertical string
		rdate    string
	}
	tests := []struct {
		name   string
		fields fields
		want   Reservation
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &reservationBuilder{
				vertical: tt.fields.vertical,
				rdate:    tt.fields.rdate,
			}
			if got := r.Build(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Build() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reservationBuilder_ReservationDate(t *testing.T) {
	type fields struct {
		vertical string
		rdate    string
	}
	type args struct {
		date string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   ReservationBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &reservationBuilder{
				vertical: tt.fields.vertical,
				rdate:    tt.fields.rdate,
			}
			if got := r.ReservationDate(tt.args.date); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReservationDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reservationBuilder_Vertical(t *testing.T) {
	type fields struct {
		vertical string
		rdate    string
	}
	type args struct {
		v string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   ReservationBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &reservationBuilder{
				vertical: tt.fields.vertical,
				rdate:    tt.fields.rdate,
			}
			if got := r.Vertical(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Vertical() = %v, want %v", got, tt.want)
			}
		})
	}
}
