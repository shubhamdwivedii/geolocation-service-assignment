Structure 


    cmd 
        - data 
            - main.go -> populate db data 
        
        - server 
            - main.go -> runs rest api service

    pkg 
        - services 
            - add 
                - service.go -> service to add geodata 
            - list 
                - service.go -> service to list (get) geodata 
        
        - http 
            - handler.go -> 
                - addGeodata(w, r) 
                - getGeodata(w, r) 
        
        - importer 
            - csv 
                - csv.go -> imports geodata from csv file or Url

        - storage 
            - sql
                - sql.go -> 
                    - init DB
                    - addGeo() to db 
                    - getGeo() from db 
    
    models 
        - geodata.go -> geodata model 