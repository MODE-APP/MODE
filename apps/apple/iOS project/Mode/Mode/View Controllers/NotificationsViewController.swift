//
//  NotificationsViewController.swift
//  Mode
//
//  Created by Jackson Tubbs on 9/17/19.
//  Copyright © 2019 Jax Tubbs. All rights reserved.
//

import UIKit

class NotificationsViewController: UIViewController {

    // MARK: - Outlets
    
    @IBOutlet weak var notificationsCollectionView: UITableView!
    
    // MARK: - Properties
    
    var loadingNotifications: Bool = false
    var loadedAllNotifications: Bool = false
    var dataSource: Int = 0
    var numberOfNotifications: Int = 60
    
    override func viewDidLoad() {
        super.viewDidLoad()
        
        dataSource = 15
        
        
        notificationsCollectionView.delegate = self
        notificationsCollectionView.dataSource = self
    }
    
    
    // MARK: - Custom Functions
    
    func getImages (completion: @escaping () -> Void) {
        let timer = Timer(timeInterval: 0.7, repeats: false) { (_) in
            completion()
        }
        RunLoop.current.add(timer, forMode: .common)
    }
    
    func loadMoreNotifications() {
        getImages() {
            DispatchQueue.main.async {
                if self.loadedAllNotifications == false {
                    self.dataSource += 15
                }
                if self.dataSource >= self.numberOfNotifications {
                    self.loadedAllNotifications = true
                }
                self.notificationsCollectionView.reloadData()
                self.loadingNotifications = false
            }
        }
    }

    /*
    // MARK: - Navigation

    // In a storyboard-based application, you will often want to do a little preparation before navigation
    override func prepare(for segue: UIStoryboardSegue, sender: Any?) {
        // Get the new view controller using segue.destination.
        // Pass the selected object to the new view controller.
    }
    */

} // End of class

extension NotificationsViewController: UITableViewDelegate, UITableViewDataSource {
    func tableView(_ tableView: UITableView, numberOfRowsInSection section: Int) -> Int {
        return dataSource
    }
    
    func tableView(_ tableView: UITableView, cellForRowAt indexPath: IndexPath) -> UITableViewCell {
        guard let cell = notificationsCollectionView.dequeueReusableCell(withIdentifier: "notificationCell", for: indexPath) as? notificationTableViewCell else {return UITableViewCell()}
        
        return cell
    }
    
    func tableView(_ tableView: UITableView, viewForFooterInSection section: Int) -> UIView? {
        
    }
    
    func tableView(_ tableView: UITableView, heightForFooterInSection section: Int) -> CGFloat {
        <#code#>
    }
}
