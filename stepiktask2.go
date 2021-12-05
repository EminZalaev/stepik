type Game struct{
     On bool
     Ammo int
     Power int
}

func (p *Game) Shoot() bool {
    if p.On==false{
        return false
    }else if p.Ammo!=0 {
     p.Ammo= p.Ammo-1
     return true
    }else{
     return false
    }
}

func (p *Game) RideBike() bool {
    if p.On==false {
        return false
    }else if p.Power!=0 {
     p.Power= p.Power-1
     return true
    }else{
     return false
    }
}
